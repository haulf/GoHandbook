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
	logFileName := "logfile"
	logFile, err := os.Create(logFileName)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("Open file error !")
	}

	debugLog = log.New(logFile, "[ai]", log.Flags()|log.Ldate|log.Lmicroseconds|log.Lshortfile)
	debugLog.Println(">>>>>>>>>>begin")

	fmt.Println("")
	fmt.Println("")
	fmt.Println("Copyright (C) aihaofeng 2017")
	fmt.Println("")
	fmt.Println("")

	outputLogInfo()

	debugLog.Println(">>>>>>>>>>end")
}

func outputLogInfo() {
	debugLog.Println("A debug message here")
}
