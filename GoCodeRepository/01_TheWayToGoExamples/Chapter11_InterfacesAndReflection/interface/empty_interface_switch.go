// file:        empty_interface_switch.go
// author:      haulf
// date:        2017.10.17
// brief:       Empty interface switch test.

package main

import (
	"fmt"
)

type specialString string

var whatIsThis specialString = "hello"

func TypeSwitch() {
	testFunc := func(any interface{}) {
		switch v := any.(type) {
		case bool:
			fmt.Printf("any %v is abool type", v)
		case int:
			fmt.Printf("any %v is anint type", v)
		case float32:
			fmt.Printf("any %v is afloat32 type", v)
		case string:
			fmt.Printf("any %v is astring type", v)
		case specialString:
			fmt.Printf("any %v is aspecial String!", v)
		default:
			fmt.Println("unknowntype!")
		}
	}
	testFunc(whatIsThis)
}

func main() {
	TypeSwitch()
}
