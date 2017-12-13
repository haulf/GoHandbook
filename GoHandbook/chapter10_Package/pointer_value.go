// @file:        pointer_value.go
// @version:     1.0
// @author:      aihaofeng
// @date:        2017.12.10
// @go version:  1.9
// @brief:       Method test.

package main

import (
	"fmt"
)

type B struct {
	thing int
}

func (b *B) change()      { b.thing = 1 }
func (b B) write() string { return fmt.Sprint(b) }

func main() {
	var b1 B // b1是值
	b1.change()
	fmt.Println(b1.write())

	b2 := new(B) // b2是指针
	b2.change()
	fmt.Println(b2.write())
}
