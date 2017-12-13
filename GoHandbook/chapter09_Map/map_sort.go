// @file:        map_sort.go
// @version:     1.0
// @author:      aihaofeng
// @date:        2017.11.19
// @go version:  1.9
// @brief:       Map sort test.

package main

import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

func main() {
	fmt.Println("\nunsorted:")

	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v \n ", k, v)
	}

	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}

	sort.Strings(keys)
	fmt.Println("\n\nsorted:")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v \n ", k, barVal[k])
	}
}
