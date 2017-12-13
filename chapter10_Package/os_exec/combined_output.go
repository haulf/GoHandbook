package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	log.SetFlags(log.Flags() | log.Ldate | log.Lmicroseconds)
	log.Println(">>>>>>>>>>begin....")
	cmd := exec.Command("sh", "-c", "echo stdout; echo 1>&2 stderr")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)
	log.Println(">>>>>>>>>>end....")
}
