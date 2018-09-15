// file:        GetFavouriteNumber.go
// author:      Haulf
// date:        2018.06.30
// brief:       Select favourite number from a lot of car numbers.

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

var debugLog *log.Logger

func main() {
	logFileName := "logfile"
	logFile, err := os.Create(logFileName)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("Open file error !")
	}

	debugLog = log.New(logFile, "", 0)

	cmdlineFile, inputError := os.Open("YueBCarNubers.txt")
	if inputError != nil {
		return
	}
	defer cmdlineFile.Close()

	inputReader := bufio.NewReader(cmdlineFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			return
		}

		resultSlice := strings.Split(inputString, ",")
		originalString := resultSlice[0]
		patternString := "16$|58$|18$|56$"
		if ok, _ := regexp.Match(patternString, []byte(originalString)); ok {
			if !strings.Contains(originalString, "4") {
				debugLog.Println(originalString)
			}
		}
	}
}
