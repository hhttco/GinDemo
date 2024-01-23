package logging

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

var (
	LogSavePath = "application/logs/"
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeFormat  = "20060102"
)

func getLogFilePath() (string, string) {
	Separator := "/"
	if runtime.GOOS == "windows" {
		Separator = "\\"
	}

	path := strings.Replace(LogSavePath, "/", Separator, -1)

	dir, _ := os.Getwd()
	dirPath := dir + Separator + path
	dirPath = strings.Replace(dirPath, "/", Separator, -1)

	return fmt.Sprintf("%s", path), fmt.Sprintf("%s", dirPath)
}

func getLogFileFullPath() string {
	prefixPath, _ := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}

	return handle
}

func mkDir() {
	_, dirPath := getLogFilePath()
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
