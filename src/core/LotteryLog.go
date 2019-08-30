package main

import (
	"fmt"
	log "logger"
	"os"
)

const (
	logPATH = "log/"
	logFile = "lottery.log"
)

/*OpenLoggerFile is for to put the logs*/
func OpenLoggerFile() (func() *os.File, error) {
	var fileHandle *os.File
	err := os.MkdirAll(logPATH, 0777)
	if err != nil {
		fmt.Println("Not able create log directory.")
		return nil, err
	}
	fileHandle, err = os.OpenFile(logPATH+logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Logger File Cannot Open. Program should not run Without logging.")
		return nil, err
	}
	log.Init(fileHandle, fileHandle, fileHandle, fileHandle)

	return func() *os.File {
		return fileHandle
	}, nil
}
