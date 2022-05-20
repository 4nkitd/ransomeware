package files

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"ransomware/exception"
	"time"
)

type FS struct {
	Name    string
	Size    int64
	IsDir   bool
	ModTime time.Time
}

func ListDir(location string) (fs []FS) {

	files, err := ioutil.ReadDir(location)

	exception.CheckErr(err)

	for _, f := range files {
		fs = append(fs, FS{
			Name:    f.Name(),
			Size:    f.Size(),
			IsDir:   f.IsDir(),
			ModTime: f.ModTime(),
		})
	}

	return fs

}

func WalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func VerifyOrCreate(location string) {
	if _, err := os.Stat(location); os.IsNotExist(err) {
		os.Mkdir(location, 0755)
	}
}
