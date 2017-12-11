/*
@file:    pointer_array.go
@version: v1.0
@author:  aihaofeng
@date:    2017.11.12
@brief:   Array test program.
*/

package main

import "fmt"

func f(a [3]int) {
	fmt.Println(a)
}

func fp(a *[3]int) {
	fmt.Println(a)
}

func main() {
	var ar [3]int
	f(ar)   // passes a copy of ar
	fp(&ar) // passes a pointer to ar
}
