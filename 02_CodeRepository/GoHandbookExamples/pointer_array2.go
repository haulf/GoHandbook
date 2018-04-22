/*
@file:    for_arrays2.go
@version: v1.0
@author:  aihaofeng
@date:    2017.11.12
@brief:   Array test program.
*/

package main

import "fmt"

func fp(a *[3]int) {
	fmt.Println(a)
}

func main() {
	for i := 0; i < 3; i++ {
		fp(&[3]int{i, i * i, i * i * i})
	}
}
