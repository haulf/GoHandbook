/*
@file:    defer_param.go
@version: v1.0
@author:  aihaofeng
@date:    2017.11.12
@brief:   Function test program.
*/

package main

import "fmt"

func main() {
	f()
}

func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("i=%d ", i)
	}
}
