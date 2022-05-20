package config

import "time"

type Configuration struct {
	LogDir  string
	LogFile string
	Key     string
}

var Data = Configuration{
	LogDir:  "./logs/",
	LogFile: time.Now().Format("2006-01-02") + ".log",
	Key:     "ankitisnotagodbutnotless",
}
