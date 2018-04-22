// file:    DeferParamSequence.go
// author:  haulf
// date:    2017.11.12
// brief:   Test the sequence of multi defer params.

package main

import "fmt"

func main() {
	f()
}

func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("i=%d ", i)
	}
}
