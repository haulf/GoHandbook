// file:        ParseCmdlineParams_New.go
// author:      haulf
// date:        2018.05.27
// brief:       Parse cmdline from kernel.
//              Reading the contents of an entire file in a string:
//              If this is sufficient for you needs, you can use the ioutil.ReadFile() methond from
//              the package io/ioutil, which returns a []byte containing the bytes read and nil or
//              a possible error. Analogously the WriteFile() function writes a []byte to a file.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	buf, err := ioutil.ReadFile("cmdline")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Read file error: %s\n", err)
	}

	resultSlice := strings.Fields(string(buf))
	for _, paramItem := range resultSlice {
		fmt.Println(paramItem)
	}
}
