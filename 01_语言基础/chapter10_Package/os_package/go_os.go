/*
@file:    go_os.go
@version: v1.0
@author:  aihaofeng
@date:    2017.08.23
@brief:   OS package test program. Get the type of operation system.
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	var goos string = os.Getenv("GOOS")
	fmt.Printf("The operating system is: \n%s\n", goos)

	path := os.Getenv("PATH")
	fmt.Printf("\nPath is:\n%s\n\n", path)
}
