/*
@file:    multiple_return.go
@version: v1.0
@author:  aihaofeng
@date:    2017.11.12
@brief:   Function test program.
*/

package main

import "fmt"

var num int = 10
var numx2, numx3 int

func main() {
	numx2, numx3 = getX2AndX3(num)
	PrintValues()
	numx2, numx3 = getX2AndX3_2(num)
	PrintValues()
}

func PrintValues() {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num,
		numx2, numx3)
}

func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	// return x2, x3
	return
}
