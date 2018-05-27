// file:    DeferTracingNew.go
// author:  haulf
// date:    2017.11.12
// brief:   Tracing function with defer.

package main

import "fmt"

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func untrace(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer untrace(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer untrace(trace("b"))
	fmt.Println("in b")
	a()
}

func main() {
	b()
}
