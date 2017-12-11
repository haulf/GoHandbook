// @file:        reflect_struct.go
// @version:     1.0
// @author:      aihaofeng
// @date:        2017.10.21
// @go version:  1.9
// @brief:       Reflect test.

package main

import (
	"fmt"
	"reflect"
)

type NotknownType struct {
	s1, s2, s3 string
}

func (n NotknownType) String() string {
	return n.s1 + " - " + n.s2 + " - " + n.s3
}

// variable to investigate:
var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}

func main() {
	value := reflect.ValueOf(secret)
	typ := reflect.TypeOf(secret)

	typeValue := value.Type()
	fmt.Println(typeValue)

	fmt.Println(typ)
	knd := value.Kind()
	fmt.Println(knd)

	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
	}

	results := value.Method(0).Call(nil)
	fmt.Println(results)
}
