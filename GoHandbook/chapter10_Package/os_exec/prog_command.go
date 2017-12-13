package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("prog")
	cmd.Env = append(os.Environ(),
		"FOO=dumplicate_value",
		"FOO=actual_value",
	)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
