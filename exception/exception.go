package exception

import (
	"ransomware/config"
	"ransomware/logger"
	"time"
)

type Exception struct {
	time     time.Time
	message  string
	location string
}

var BackLog []Exception

func Handle(err error) (e Exception) {

	log := logger.NewLogger(config.Data.LogDir+config.Data.LogFile, true)
	log.Debug("Error Occurred: " + err.Error())
	var handler = Handler{}

	handler.Handle(err)

	BackLog = append(BackLog, handler.ParseError())

	return e
}

func CheckErr(err error) {
	if err != nil {
		Handle(err)
	}
}

func RunNotNull(val any) (e bool) {
	if val != nil {
		return true
	}
	return true
}
