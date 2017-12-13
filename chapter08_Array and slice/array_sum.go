/*
@file:    array_sum.go
@version: v1.0
@author:  aihaofeng
@date:    2017.11.12
@brief:   Array test program.
*/

package main

import "fmt"

func main() {
	array := [3]float64{7.0, 8.5, 9.1}
	x := Sum(&array) // Note the explicit address-of operator

	// to pass a pointer to the array
	fmt.Printf("The sum of the array is: %f", x)
}

func Sum(a *[3]float64) (sum float64) {
	// derefencing *a to get back to the array is not necessary!
	for _, v := range a {
		sum += v
	}

	return
}
