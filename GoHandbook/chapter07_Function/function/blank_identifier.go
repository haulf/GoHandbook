/*
@file:    blank_identifier.go
@version: v1.0
@author:  aihaofeng
@date:    2017.11.12
@brief:   Function test program.
*/

package main

import "fmt"

func main() {
	var i1 int
	var f1 float32
	i1, _, f1 = getThreeValues()
	fmt.Printf("The int: %d, the float: %f \n", i1, f1)
}

func getThreeValues() (int, int, float32) {
	return 5, 6, 7.5
}
