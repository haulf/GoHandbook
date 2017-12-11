// @file:        channel_block.go
// @version:     1.0
// @author:      aihaofeng
// @date:        2017.11.08
// @go version:  1.9
// @brief:       Goroutine test.

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go pump(ch1) // pump hangs
	go suck(ch1)
	time.Sleep(1e9)
	fmt.Println(<-ch1) // prints only 0
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
