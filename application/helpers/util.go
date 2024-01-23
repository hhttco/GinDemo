package helpers

import (
	"os"
	"runtime"
)

/**
 * 判断文件夹是否存在
 *
 */
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/**
 * 获取本地路径
 *
 */
func GetProjectPath() string {
	var projectPath string
	projectPath, _ = os.Getwd()
	return projectPath
}

/**
 * 判断当前操作系统
 *
 */
var osType = runtime.GOOS

func GetOsType() string {
	return osType
}

func GetPathSeparator() string {
	if osType == "windows" {
		return "\\"
	} else {
		return "/"
	}
}
