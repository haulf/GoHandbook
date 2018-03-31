// @file:        File write.go
// @author:      haulf
// @date:        2018.02.20
// @go version:  1.9
// @brief:       File write test.

package main

import "os"

func main() {
	os.Stdout.WriteString("hello, world\n")
	f, _ := os.OpenFile("test_file", os.O_CREATE|os.O_WRONLY, 0)
	defer f.Close()

	f.WriteString("hello, world in a file\n")
}
