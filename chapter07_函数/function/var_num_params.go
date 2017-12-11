/*
@file:    var_num_params.go
@version: v1.0
@author:  aihaofeng
@date:    2017.11.12
@brief:   Function test program.
*/

package main

import "fmt"

func main() {
	x := min(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)

	arr := []int{7, 9, 3, 5, 1}
	x = min(arr...)
	fmt.Printf("The minimum in the array arr is: %d", x)
}

func min(a ...int) (minValue int) {
	if len(a) == 0 {
		return 0
	}

	minValue = a[0]
	for _, v := range a {
		if v < minValue {
			minValue = v
		}
	}

	return
}
