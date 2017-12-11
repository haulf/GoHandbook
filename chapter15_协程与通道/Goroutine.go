// author:     aihaofeng
// date:       2017.08.22
// version:    1.0
// go version: 1.9
// brief:      Goroutine test.

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("Hello")
	say("aihaofeng")
}
