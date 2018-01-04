/*
@file:    interface_poly.go
@version: v1.0
@author:  aihaofeng
@date:    2017.10.17
@brief:   Interface test program.
*/

package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func main() {
	r := Rectangle{5, 3}
	q := &Square{5}
	shapers := []Shaper{Shaper(r), Shaper(q)}
	// shapers := []Shaper{r, q}

	fmt.Println("Looping through shapes for area ...")

	for n, _ := range shapers {
		fmt.Println("Shape details:", shapers[n])
		fmt.Println("Area of this shape is:", shapers[n].Area())
	}
}
