// @file:        structs_fields.go
// @version:     1.0
// @author:      aihaofeng
// @date:        2017.12.10
// @go version:  1.9
// @brief:       Struct test.

package main

import (
	"fmt"
)

type TestStruct struct {
	iField   int
	fField   float32
	strField string
}

func main() {
	ts := new(TestStruct)
	ts.iField = 10
	ts.fField = 15.5
	ts.strField = "aihaofeng"

	fmt.Printf("The int is: %d\n", ts.iField)
	fmt.Printf("The float is: %f\n", ts.fField)
	fmt.Printf("The string is: %s\n", ts.strField)
	fmt.Println(ts)
}
