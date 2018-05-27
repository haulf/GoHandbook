/*
@file:    return_defer.go
@version: v1.0
@author:  aihaofeng
@date:    2017.11.12
@brief:   Function test program.
*/

package main

import "fmt"

func f() (ret int) {
	defer func() {
		ret++
	}()

	return 1
}

func main() {
	fmt.Println(f())
}
