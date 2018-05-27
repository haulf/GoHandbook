// file:    RecordParametersAndReturns.go
// author:  aihaofeng
// date:    2017.11.12
// brief:   Record function's parameters and returns by defer.

package main

import (
	"io"
	"log"
)

func main() {
	myfunction("Go")
}

func myfunction(s string) (n int, err error) {
	defer func() {
		log.Printf("myfunction(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}
