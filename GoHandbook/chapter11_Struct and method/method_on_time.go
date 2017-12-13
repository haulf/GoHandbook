// @file:        method_on_time.go
// @version:     1.0
// @author:      aihaofeng
// @date:        2017.12.10
// @go version:  1.9
// @brief:       Method test.

package main

import (
	"fmt"
	"time"
)

type myTime struct {
	time.Time // anonymous field
}

func (t myTime) first3Chars() string {
	return t.Time.String()[0:3]
}

func main() {
	m := myTime{time.Now()}
	// 调用匿名Time上的String方法
	fmt.Println("Full time now:", m.String())
	// 调用myTime.first3Chars
	fmt.Println("First 3 chars:", m.first3Chars())
}
