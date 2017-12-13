// author:     aihaofeng
// date:       2017.08.22
// version:    1.0
// go version: 1.9
// brief:      Goroutine, channel, and time test.

package main

import (
	"fmt"
	"time"
)

// 为了实现goroutine，准备一个函数，用通道作为参数。
func getNowTime(ch chan string) {
	nowTime := time.Now().Format("2006年01月02日 15点04分05秒.0000000 时区-0700")
	ch <- nowTime
}

func main() {
	// 创建一个通道。
	ch := make(chan string)

	for i := 0; i < 10; i++ {
		go getNowTime(ch)
		fmt.Println(<-ch)

		// 每次循环间隔0.5秒。
		time.Sleep(500 * time.Millisecond)
	}
}
