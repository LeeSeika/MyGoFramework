package util

import "os"

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
