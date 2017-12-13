package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	path, err := exec.LookPath("fortune")
	if err != nil {
		log.Fatal("Install fortune is in the future.")
	}
	fmt.Println("Fortune is available at %s\n", path)
}
