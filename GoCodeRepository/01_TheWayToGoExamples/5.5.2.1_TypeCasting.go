// file:    5.5.2.1_TypeCasting.go
// author:  haulf
// date:    2017.08.27
// brief:   Type casting test program.

package main

import (
	"fmt"
)

func main() {
	var n int16 = 34
	var m int32
	// compiler error: cannot use n (type int16) as type int32 in assignment
	//m = n
	m = int32(n)

	fmt.Printf("32 bit int is: %d\n", m)
	fmt.Printf("16 bit int is: %d\n", n)
}
