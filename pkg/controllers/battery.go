package controllers

import (
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/byeol-i/battery-level-checker/pkg/device"
	"github.com/byeol-i/battery-level-checker/pkg/logger"
	"github.com/byeol-i/battery-level-checker/pkg/producer"
	dbSvc "github.com/byeol-i/battery-level-checker/pkg/svc/db"
	"go.uber.org/zap"
	// "github.com/byeol-i/battery-level-checker/pkg/models"
)

type BatteryController struct {
	basePattern string
}

func NewBatteryController(basePattern string) *BatteryController {
	return &BatteryController{
		basePattern: basePattern,
	}
}

// GetUsersAllBattery godoc
// @Summary Get User's all Battery Level
// @Description Get device's battery
// @Tags Battery
// @Accept json
// @Produce json
// @Param Authorization header string true "With the bearer started"
// @Failure 400 {object} models.JSONfailResult{}
// @Success 200 {object} models.JSONsuccessResult{data=[]device.DeviceImpl{Id,BatteryLevel}}
// @Router /battery/ [get]
func (hdl *BatteryController) GetUsersAllBattery(resp http.ResponseWriter, req *http.Request) {
	uid := req.Header.Get("Uid")



	// test

	// err := consumer.ConsumeLatestMessage("test_test")
	// if err != nil {
	// 	logger.Error("Can't consume msg", zap.Error(err))
	// }

	res, err := dbSvc.CallGetUsersAllBattery(uid)
	if err != nil {
		logger.Error("dbSvc's error", zap.Error(err))
		respondError(resp, http.StatusInternalServerError, "Internal server error")
		return
	}

	// consumer.GetTopics()
	respondJSON(resp, http.StatusOK, "GetUsersAllBattery", res)
}


// GetSpecificDevice's battery godoc
// @Summary Get Specific Device's Battery Level
// @Description Get device battery
// @Tags Battery
// @Accept json
// @Produce json
// @Param id path string true "Device ID"
// @Param Authorization header string true "With the bearer started"
// @Failure 400 {object} models.JSONfailResult{}
// @Success 200 {object} models.JSONsuccessResult{data=device.DeviceImpl{Id,BatteryLevel}}
// @Router /battery/ [get]
func (hdl *BatteryController) GetBattery(resp http.ResponseWriter, req *http.Request) {
	pattern := regexp.MustCompile(hdl.basePattern + `/battery/([a-zA-Z0-9-]+)`)
	
    matches := pattern.FindStringSubmatch(req.URL.Path)
	if len(matches) == 0 {
        respondError(resp, http.StatusBadRequest, "Not valid")
        return
    }

	// uid := "P1RPAH3URaYL0FoZdjLKJf20h9T2"
	
	// uid := req.Header.Get("Uid")

	// res, err := dbSvc.CallGetBattery(matches[1], uid)
	// if err != nil  {
	// 	logger.Error("dbSvc's error", zap.Error(err))
	// 	respondError(resp, http.StatusInternalServerError, "Internal server error")
	// 	return
	// }

	// err := consumer.ConsumeLatestMessage("battery_device__"+uid+"_"+matches[1])
	// if err != nil {
	// 	logger.Error("Can't consume msg", zap.Error(err))
	// }

	// consumer.GetTopics()
	// respondJSON(resp, http.StatusOK, "GetBattery", res)

	respondJSON(resp, http.StatusOK, "GetBattery", nil)
}



// GetHistoryAllBattery godoc
// @Summary Get All Device's Battery Level
// @Description Get devices's battery
// @Tags Battery
// @Accept json
// @Produce json
// @Param Authorization header string true "With the bearer started"
// @Failure 400 {object} models.JSONfailResult{}
// @Success 200 {object} models.JSONsuccessResult{data=[]device.DeviceImpl{Id,BatteryLevel}}
// @Router /battery/history/{deviceId}} [get]
func (hdl *BatteryController) GetHistoryAllBattery(resp http.ResponseWriter, req *http.Request) {
	pattern := regexp.MustCompile(`/battery/history/(\w+)`)
    matches := pattern.FindStringSubmatch(req.URL.Path)
	if len(matches[1]) < 2 {
        respondError(resp, http.StatusBadRequest, "Not valid")
        return
    }

	uid := req.Header.Get("Uid")
	logger.Info("GetHistoryAllBattery", zap.String("uid", uid), zap.String("device", matches[1]))
	
	res, err := dbSvc.CallGetAllBattery(matches[1], uid)
	if err != nil {
		logger.Error("dbSvc's error", zap.Error(err))
		respondError(resp, http.StatusInternalServerError, "Internal server error")
		return
	}

	// consumer.GetTopics()
	respondJSON(resp, http.StatusOK, "GetHistoryAllBattery", res)
}

// UpdateBatteryLevel godoc
// @Summary Update Battery Level
// @Description Update devices's battery
// @Tags Battery
// @Accept json
// @Produce json
// @Param device body device.DeviceImpl true "Battery input form"
// @Param Authorization header string true "With the bearer started"
// @Success 200 {object} models.JSONsuccessResult{}
// @Failure 400 {object} models.JSONfailResult{}
// @Router /battery/{deviceID} [put]
func (hdl *BatteryController) UpdateBattery(resp http.ResponseWriter, req *http.Request) {
	pattern := regexp.MustCompile(hdl.basePattern + `/battery/([a-zA-Z0-9-]+)`)
	
    deviceId := pattern.FindStringSubmatch(req.URL.Path)
	if len(deviceId) < 2 {
        respondError(resp, http.StatusBadRequest, "Not valid")
        return
    }
	
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Error("Can't reading body", zap.Error(err))
	}

	newBatteryLevel, err := device.ParseFromJson(string(body[:]))
	if err != nil {
		logger.Error("Json parse error", zap.Error(err))
		respondError(resp, http.StatusRequestedRangeNotSatisfiable, "invalid format")
		return
	}

	err = device.BatteryLevelValidator(newBatteryLevel)
	if err != nil {
		logger.Error("Battery level is not validate", zap.Error(err))	
		respondError(resp, http.StatusRequestedRangeNotSatisfiable, "invalid format")
		return
	}

	

	uid := "P1RPAH3URaYL0FoZdjLKJf20h9T2"
	
	// uid := req.Header.Get("Uid")

	// err = dbSvc.CallUpdateBatteryLevel(deviceId[1], uid, newBatteryLevel)
	// if err != nil {
	// 	logger.Error("dbSvc's error", zap.Error(err))
	// 	respondError(resp, http.StatusInternalServerError, "Internal server error")
	// 	return
	// }

	err = producer.WriteBatteryTime(newBatteryLevel, deviceId[1], uid)
	if err != nil {
		logger.Error("can't producer ", zap.Error(err))
		respondError(resp, http.StatusInternalServerError, "Internal server error")
		return
	}
	// t, err := time.Parse("2006-01-02 15:04:05", req.PostFormValue("Time"))
	// if err != nil {
	// 	respondError(resp, 405, "time error")
	// }

	// bt, err := strconv.Atoi(req.PostFormValue("BatteryLevel"))
	// if err != nil {
	// 	respondError(resp, 405, "can't convert batteryLevel")
	// }

	// validDevice := device.DeviceImpl{
	// 	Name : req.PostFormValue("Name"),
	// 	Time : &t,
	// 	BatteryLevel: bt,
	// 	BatteryStatus: req.PostFormValue("BatteryStatus"),
	// }

	// v := validator.New()

	// err = v.Struct(validDevice)
	// if err != nil {
	// 	respondError(resp, 405, "not valid form")
	// }

	// producer.Write() 

	respondJSON(resp, 200, "UpdateBattery", "")
}
