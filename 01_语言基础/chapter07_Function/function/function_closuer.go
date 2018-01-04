// @file:        function_closure.go
// @version:     1.0
// @author:      aihaofeng
// @date:        2017.11.05
// @go version:  1.9
// @brief:       Function return test.

package main

import "fmt"

func main() {
	var f = Adder()
	fmt.Print("")
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))
	fmt.Print("")
}

func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
