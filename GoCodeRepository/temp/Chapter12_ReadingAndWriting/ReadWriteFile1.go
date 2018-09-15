// @file:        ReadWriteFile1.go
// @author:      haulf
// @date:        2017.12.25
// @brief:       Read and write test.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := "products.txt"
	outputFile := "products_copy.txt"

	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		// panic(err.Error())
	}

	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644) // oct, not hex
	if err != nil {
		panic(err.Error())
	}
}
