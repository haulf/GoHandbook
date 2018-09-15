// @file:        copy_append_slice.go
// @version:     1.0
// @author:      aihaofeng
// @date:        2017.11.15
// @go version:  1.9
// @brief:       Slice test.

package main

import (
	"fmt"
)

func main() {
	sliceFrom := []int{1, 2, 3}
	sliceTo := make([]int, 10)
	n := copy(sliceTo, sliceFrom)
	fmt.Println(sliceTo)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sliceNew := []int{1, 2, 3}
	sliceNew = append(sliceNew, 4, 5, 6)
	fmt.Println(sliceNew)
}
