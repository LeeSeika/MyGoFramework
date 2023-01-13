package env

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/godemo/coredemo/framework/contract"
	"io"
	"os"
	"strings"
)

type HadeEnv struct {
	envMap map[string]string
	folder string
}

func (he *HadeEnv) AppEnv() string {
	if value, ok := he.envMap["APP_ENV"]; ok {
		return value
	}
	return ""
}

func (he *HadeEnv) IsExist(key string) bool {
	_, ok := he.envMap[key]
	return ok
}

func (he *HadeEnv) Get(key string) string {
	if value, ok := he.envMap[key]; ok {
		return value
	}
	return ""
}

func (he *HadeEnv) All() map[string]string {
	return he.envMap
}

func NewHadeEnvService(params ...interface{}) (interface{}, error) {
	if len := len(params); len != 1 {
		fmt.Println("error params")
		return nil, errors.New("params error")
	}
	folder := params[0].(string)
	hadeEnv := HadeEnv{
		envMap: map[string]string{"APP_ENV": contract.EnvDevelopment},
		folder: folder,
	}
	file, err := os.Open(folder)
	if err == nil {
		reader := bufio.NewReader(file)
		for {
			// 按行读
			line, _, err := reader.ReadLine()
			// 死循环通过err == EOF结束
			if err == io.EOF {
				break
			}
			// 将行的byte数据分割
			splitN := bytes.SplitN(line, []byte{'='}, 2)
			// 不符合要求的直接无视
			if len(splitN) < 2 {
				continue
			}
			key := string(splitN[0])
			value := string(splitN[1])
			hadeEnv.envMap[key] = value
		}
		environ := os.Environ()
		for _, line := range environ {
			splitN := strings.SplitN(line, "=", 2)
			if len(splitN) < 2 {
				continue
			}
			key := splitN[0]
			value := splitN[1]
			hadeEnv.envMap[key] = value
		}
	}
	return &hadeEnv, nil
}
