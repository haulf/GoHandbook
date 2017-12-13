// @file:        channel_block.go
// @version:     1.0
// @author:      aihaofeng
// @date:        2017.11.08
// @go version:  1.9
// @brief:       Goroutine test.

package main

import "fmt"

func main() {
	ch1 := make(chan int)
	go pump(ch1)       // pump hangs
	fmt.Println(<-ch1) // prints only 0
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}
