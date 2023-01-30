package util

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"syscall"
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

// Unzip 解压缩zip文件，复制文件和目录都到目标目录中
func Unzip(src string, dest string) ([]string, error) {

	var filenames []string

	// 使用archive/zip读取
	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	// 所有内部文件都读取
	for _, f := range r.File {

		// 目标路径
		fpath := filepath.Join(dest, f.Name)

		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// 如果是目录，则创建目录
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		//否则创建文件
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		// 复制内容
		_, err = io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}

func CheckProcessExist(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	err = process.Signal(syscall.Signal(0))
	if err != nil {
		return false
	}
	return true
}
