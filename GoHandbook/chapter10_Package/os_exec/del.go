// author:     aihaofeng
// date:       2017.08.22
// version:    1.0
// go version: 1.9
// brief:      Del file test.

package main

import (
	"fmt"
	"os/exec"
)

func main() {
	execCommand()
}

func execCommand() {
	cmd := exec.Command("CMD", "/C", "del", "D:\\system_zip.img")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
