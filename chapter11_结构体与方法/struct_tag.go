// @file:        struct_tag.go
// @version:     1.0
// @author:      aihaofeng
// @date:        2017.12.10
// @go version:  1.9
// @brief:       Struct test.

package main

import (
	"fmt"
	"reflect"
)

type TagType struct { // tags
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much thereare"
}

func main() {
	tt := TagType{true, "Barak Obama", 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)

	fmt.Printf("%v\n", ixField.Tag)
}
