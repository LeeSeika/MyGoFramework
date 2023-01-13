package services

import (
	"errors"
	"fmt"
	"github.com/godemo/coredemo/framework"
	"github.com/godemo/coredemo/framework/contract"
	"github.com/godemo/coredemo/framework/util"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"os"
	"path/filepath"
	"time"
)

func NewHadeRotateLog(params ...interface{}) (interface{}, error) {
	if len(params) != 4 {
		return nil, errors.New("params error")
	}
	container := params[0].(framework.Container)
	level := params[1].(contract.LogLevel)
	ctxFielder := params[2].(contract.CtxFielder)
	formatter := params[3].(contract.Formatter)

	rotateLog := &HadeRotateLog{}

	appService := container.MustMake(contract.AppKey).(contract.App)
	configService := container.MustMake(contract.ConfigKey).(contract.Config)
	// folder
	folder := appService.LogFolder()
	if configService.IsExist("log.folder") {
		configService.GetString("log.folder")
		rotateLog.folder = folder
	}
	if !util.Exists(rotateLog.folder) {
		err := os.MkdirAll(rotateLog.folder, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	// fileName
	fileName := "hade.log"
	if configService.IsExist("log.file") {
		fileName = configService.GetString("log.file")
	}

	// date
	dateFormat := "%Y%m%d%H"
	if configService.IsExist("log.date_format") {
		dateFormat = configService.GetString("log.date_format")
	}

	fullPath := rotatelogs.WithLinkName(filepath.Join(folder, fileName))
	options := []rotatelogs.Option{fullPath}

	// rotate_count
	if configService.IsExist("log.rotate_count") {
		rotateCount := configService.GetInt("log.rotate_count")
		options = append(options, rotatelogs.WithRotationCount(uint(rotateCount)))
	}
	// rotate_size
	if configService.IsExist("log.rotate_size") {
		rotateSize := configService.GetInt("log.rotate_size")
		options = append(options, rotatelogs.WithRotationSize(int64(rotateSize)))
	}
	// rotate_time
	if configService.IsExist("log.rotate_time") {
		rotateTime := configService.GetString("log.rotate_time")
		if timeDuration, err := time.ParseDuration(rotateTime); err != nil {
			options = append(options, rotatelogs.WithRotationTime(timeDuration))
		}
	}
	// max_age
	if configService.IsExist("log.max_age") {
		maxAge := configService.GetString("log.max_age")
		if timeDuration, err := time.ParseDuration(maxAge); err != nil {
			options = append(options, rotatelogs.WithMaxAge(timeDuration))
		}
	}
	// basic info
	rotateLog.SetFormatter(formatter)
	rotateLog.SetLevel(level)
	rotateLog.SetCtxFielder(ctxFielder)
	rotateLog.fileName = fileName
	rotateLog.folder = folder
	rotateLog.container = container

	output, err := rotatelogs.New(fmt.Sprintf("%s.%s", filepath.Join(rotateLog.folder, rotateLog.fileName), dateFormat), options...)
	if err != nil {
		return nil, err
	}

	rotateLog.SetOutput(output)

	return rotateLog, nil
}

type HadeRotateLog struct {
	HadeLog
	folder   string
	fileName string
}
