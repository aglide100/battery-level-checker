package config

import (
	"flag"
	"os"

	"github.com/byeol-i/battery-level-checker/pkg/logger"
)


var (
	local = flag.Bool("localcred",  false, "using local cred")
	credFilePath = flag.String("firebaseCredFilePath", "/run/secrets/firebase-key", "cred path")
	firebaseProjectID = flag.String("firebaseProjectID", "worker-51312", "firebaseProjectID")
)

func GetFirebaseCredFilePath() string {
	flag.Parse()

	if *local {
		logger.Info("Using local conf")
		return "conf/firebase/key.json"
	}

	return *credFilePath
}

func GetFirebaseProjectID() string {
	flag.Parse()
	return *firebaseProjectID
}

func GetApiKey() string {
	key := os.Getenv("API_KEY")

	return key
}

func GetAppID() string {
	appID := os.Getenv("APP_ID")

	return appID
}
