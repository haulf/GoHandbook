package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	for i := 0; i < 10; i++ {
		go timeNow(ch)
		fmt.Println(<-ch)
		time.Sleep(500 * time.Millisecond)
	}
}

func timeNow(ch chan string) {
	tn := time.Now().Format("2018年04月21日 20点08分29秒.0000000 时区-0700")
	ch <- tn
}
