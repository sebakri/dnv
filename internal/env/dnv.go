package env

import "os"

type DNV struct {
	SessionId     string
	SessionFolder string
	EnvLoaded     string
	Shell         string
	Debug         bool
}

var dnvEnv DNV

func init() {
	dnvEnv.SessionId = os.Getenv("DNV_SESSION_ID")
	dnvEnv.SessionFolder = os.Getenv("DNV_SESSION_FOLDER")
	dnvEnv.EnvLoaded = os.Getenv("DNV_ENV_LOADED")
	dnvEnv.Shell = os.Getenv("DNV_SHELL")
	dnvEnv.Debug = os.Getenv("DNV_DEBUG") == "true"
}

func GetDNV() DNV {
	return dnvEnv
}
