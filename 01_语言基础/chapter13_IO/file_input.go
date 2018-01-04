// @file:        Demo.go
// @author:      haulf
// @date:        2017.12.19
// @go version:  1.9
// @brief:       File input test.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, inputError := os.Open("input.dat")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces toit?\n")
		return // exit the function onerror
	}

	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			return
		}

		fmt.Printf("The input was: %s", inputString)
	}
}
