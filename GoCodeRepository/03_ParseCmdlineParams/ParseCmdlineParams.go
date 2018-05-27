// file:        ParseCmdlineParameters.go
// author:      haulf
// date:        2018.02.20
// brief:       Parse cmdline from kernel.

package main

import (
    "os"
    "bufio"
    "io"
    "strings"
    "fmt"
)

func main() {
    cmdlineFile, inputError := os.Open("cmdline")
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

        resultSlice := strings.Fields(inputString)
        for _, item := range resultSlice {
            fmt.Println(item)
        }
    }
}


