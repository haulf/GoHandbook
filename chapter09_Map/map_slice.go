// @file:        map_slice.go
// @version:     1.0
// @author:      aihaofeng
// @date:        2017.11.17
// @go version:  1.9
// @brief:       Map test.

package main

import (
	"fmt"
)

func main() {
	// Version A:
	items := make([]map[int]int, 5) // 切片的元素类型是map[int]int
	for i := range items {          // items是一个切片
		items[i] = make(map[int]int, 1) // items[i]是一个Map
		items[i][1] = 2
	}

	fmt.Printf("Version A: Value of items: %v\n", items)

	// Version B: NOT GOOD!
	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		// item is only a copy of the slice element.
		item = make(map[int]int, 1)
		item[1] = 2 // This 'item' will be lost on the next iteration.
	}

	fmt.Printf("Version B: Value of items: %v\n", items2)
}
