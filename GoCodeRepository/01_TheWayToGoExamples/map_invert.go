// @file:        map_invert.go
// @version:     1.0
// @author:      aihaofeng
// @date:        2017.11.19
// @go version:  1.9
// @brief:       Map invert test.

package main

import (
	"fmt"
)

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

func main() {
	invMap := make(map[int]string, len(barVal))
	for k, v := range barVal {
		invMap[v] = k
	}

	fmt.Println("\ninverted:\n")
	for k, v := range invMap {
		fmt.Printf("Key: %v, Value: %v \n ", k, v)
	}
}
