package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"github.com/byeol-i/battery-level-checker/pkg/logger"
	dbSvc "github.com/byeol-i/battery-level-checker/pkg/svc/db"
	firebaseSvc "github.com/byeol-i/battery-level-checker/pkg/svc/firebase"
	"go.uber.org/zap"

	"github.com/byeol-i/battery-level-checker/pkg/user"
)

type UserControllers struct {
	basePattern string
}

func NewUserController(basePattern string) *UserControllers {
	return &UserControllers{
		basePattern: basePattern,
	}
}

// AddUser godoc
// @Summary Register New user
// @Description Register New user
// @Tags user
// @Accept json
// @Produce json
// @Param userInfo body user.UserImpl true "add user"
// @Param Authorization header string true "With the bearer started"
// @Failure 400 {object} models.JSONfailResult{}
// @Success 200 {object} models.JSONsuccessResult{}
// @Router /user/register [post]
func (hdl *UserControllers) AddNewUser(resp http.ResponseWriter, req *http.Request) {
	var userSpec user.UserImpl

	token := req.Header.Get("Authorization")

	if len(token) < 5 {
		respondError(resp, 401, "Can't find token")
	}
	var userCredential user.UserCredential

	const bearerPrefix = "Bearer "
	if strings.HasPrefix(token, bearerPrefix) {
		token = token[len(bearerPrefix):]

		uid, err := firebaseSvc.CallVerifyToken(token)
		if err != nil {
			logger.Error("get some error", zap.Error(err))
			respondError(resp, 401, err.Error())
		}
		userCredential.Uid = uid
	}

	err := json.NewDecoder(req.Body).Decode(&userSpec)
	if err != nil {
		respondError(resp, http.StatusBadRequest, "invalid format")
		return
	}

	err = user.UserValidator(&userSpec)
	if err != nil {
		respondError(resp, http.StatusBadRequest, err.Error())	
		return
	}

	newUser := user.NewUser()
	newUser.SetUserCredential(userCredential)
	newUser.SetName(userSpec.Name)

	err = dbSvc.CallAddNewUser(&newUser.UserImpl, &userCredential)
	if err != nil {
        respondError(resp, http.StatusBadRequest, "User's form is not valid")
		return
	}

	respondJSON(resp, http.StatusOK, "success", nil)
}

// CreateCustomToken godoc
// @Summary Create custom token
// @Description Create custom token
// @Tags user
// @Accept json
// @Produce json
// @Param Authorization header string true "With the bearer started"
// @Failure 400 {object} models.JSONfailResult{}
// @Success 200 {object} models.JSONsuccessResult{}
// @Router /user/token [post]
func (hdl *UserControllers) CreateCustomToken(resp http.ResponseWriter, req *http.Request) {
	var tokenSpec user.Token

	tokenSpec.Token = req.Header.Get("Authorization")
	tokenSpec.Uid = req.Header.Get("Uid")
	// err := json.NewDecoder(req.Body).Decode(&tokenSpec)
	// if err != nil {
	// 	respondError(resp, http.StatusBadRequest, "invalid format")
	// 	return
	// }

	err := user.TokenValidator(&tokenSpec)
	if err != nil {
		respondError(resp, http.StatusBadRequest, err.Error())	
		return
	}

	customToken, err := firebaseSvc.CallCreateCustomToken(tokenSpec)
	if err != nil {
        respondError(resp, http.StatusBadRequest, err.Error())
		return
	}

	respondJSON(resp, http.StatusOK, "success", customToken)
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete user
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "user ID"
// @Param Authorization header string true "With the bearer started"
// @Failure 400 {object} models.JSONfailResult{}
// @Success 200 {object} models.JSONsuccessResult{}
// @Router /user/{userID} [delete]
func (hdl *UserControllers) DeleteUser(resp http.ResponseWriter, req *http.Request) {
	pattern := regexp.MustCompile(hdl.basePattern + `/user/([A-Za-z])$`)
    matches := pattern.FindStringSubmatch(req.URL.Path)
	if len(matches) < 2 {
        respondError(resp, http.StatusBadRequest, "Not valid in Delete User")
        return
    }

	respondJSON(resp, http.StatusOK, "success", nil)
}