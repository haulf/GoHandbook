/*
@file:    slice_param.go
@version: v1.0
@author:  aihaofeng
@date:    2017.11.12
@brief:   Array test program.
*/

package main

import (
	"fmt"
)

func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

func main() {
	var arr = [5]int{0, 1, 2, 3, 4}
	fmt.Println("The sum is:", sum(arr[:]))
}
