// author:     aihaofeng
// date:       2017.08.22
// version:    1.0
// go version: 1.9
// brief:      Log output test.

package main

import (
	"fmt"
	"log"
	"os"
)

var debugLog *log.Logger

func main() {
	initLog()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Copyright (C) aihaofeng 2017")
	fmt.Println("")
	fmt.Println("")

	debugLog.Println("main:A debug message here")
}

func initLog() {
	logFileName := "logfile"
	logFile, err := os.Create(logFileName)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("Open file error !")
	}

	debugLog = log.New(logFile, "[aihaofeng]", log.LstdFlags)
	debugLog.Println("init:A debug message here")
}
