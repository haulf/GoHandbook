// @file:        for_range.go
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
	s := "\u00ff\u754c"
	for i, c := range s {
		fmt.Printf("%d:%c ", i, c)
	}
}
