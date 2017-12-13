// @file:        map_func.go
// @version:     1.0
// @author:      aihaofeng
// @date:        2017.09.08
// @go version:  1.9
// @brief:       Demo test.

package main

import (
	"fmt"
)

func main() {
	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		5: func() int { return 50 },
	}

	fmt.Println(mf)
}
