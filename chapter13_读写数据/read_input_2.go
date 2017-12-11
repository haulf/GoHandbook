// @file:        read_input_2.go
// @version:     1.0
// @author:      aihaofeng
// @date:        2017.12.11
// @go version:  1.9
// @brief:       Read test.

package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader
var input string
var err error

func main() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}
}
