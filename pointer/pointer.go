// @file:    5.9_pointer.go
// @version: v1.0
// @author:  aihaofeng
// @date:    2017.08.13
// @brief:   Pointer test program. Also introduce two methods of converting integer to string.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, j := 45, 1973
	p := &i
	fmt.Println("The value of *p is:" + strconv.Itoa(*p))
	*p = 27
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println("The value of j is:" + fmt.Sprintf("%d", j))
}
