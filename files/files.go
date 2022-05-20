package files

import (
	"io/ioutil"
	"os"
	"ransomware/exception"
)

type file struct{}

func Open(location string) (c string) {

	content, err := ioutil.ReadFile(location)
	exception.CheckErr(err)
	c = string(content)

	return c
}

func Write(file, content string) {
	fi, err := os.Open(file)
	exception.CheckErr(err)
	defer fi.Close()
	fi.Truncate(0) // comment or uncomment
	fi.Seek(0, 0)
	fi.WriteString((content))
}
