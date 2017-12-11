// @file:        MyStruct.go
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
	a int
	b int
}

type B struct {
	a int
	b int
}

type C struct {
	A
	B
}

func main() {
	var c C
	c.A.a = 5
	c.A.b = 6
	c.B.a = 7
	c.B.b = 8

	// error
	fmt.Println(c.a)

	// OK
	fmt.Println(c.A.a)
	fmt.Println(c.A.b)
	fmt.Println(c.B.a)
	fmt.Println(c.B.b)
}
