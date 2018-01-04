// @file:        reflect2.go
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

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)

	// setting a value:
	// Error: reflect.Value.SetFloat using unaddressable value
	// v.SetFloat(3.1415)

	fmt.Println("settability of v:", v.CanSet())

	v = reflect.ValueOf(&x) //Note: take the address of x.
	fmt.Println("type of v:", v.Type())
	fmt.Println("settability of v:", v.CanSet())

	v = v.Elem()
	fmt.Println("The Elem of v is: ", v)
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(3.1415) // this works!
	fmt.Println(v.Interface())
	fmt.Println(v)
}
