/*
@file:    side_effect.go
@version: v1.0
@author:  aihaofeng
@date:    2017.11.12
@brief:   Function test program.
*/

package main

import (
	"fmt"
)

// this function changes reply
func Multiply(a, b int, reply *int) {
	*reply = a * b
}

func main() {
	n := 0
	reply := &n
	Multiply(10, 5, reply)
	fmt.Println("Multiply:", *reply) // Multiply: 50
}
