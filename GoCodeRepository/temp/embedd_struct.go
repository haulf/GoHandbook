// @file:        embedd_struct.go
// @version:     1.0
// @author:      aihaofeng
// @date:        2017.12.10
// @go version:  1.9
// @brief:       Struct test.

package main

import (
	"fmt"
)

type A struct {
	ax, ay int
}

type B struct {
	A
	bx, by float32
}

func main() {
	b := B{A{1, 2}, 3.0, 4.0}

	fmt.Println(b.ax, b.ay, b.bx, b.by)
	fmt.Println(b.A)
}
