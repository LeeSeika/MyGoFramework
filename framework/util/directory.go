package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func GetExecDirectory() string {
	dir, err := os.Getwd()
	if err == nil {
		return dir + "/"
	}
	return ""
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func IsHiddenDirectory(path string) bool {
	return len(path) > 1 && strings.HasPrefix(filepath.Base(path), ".")
}

func SubDir(folder string) ([]string, error) {
	dir, err := ioutil.ReadDir(folder)
	if err != nil {
		return nil, err
	}
	var subFolders []string
	for _, subFolder := range dir {
		if subFolder.IsDir() {
			subFolders = append(subFolders, subFolder.Name())
		}
	}
	return subFolders, nil
}
