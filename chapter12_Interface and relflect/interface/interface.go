/*
@file:    interface_demo.go
@version: v1.0
@author:  aihaofeng
@date:    2017.09.28
@brief:   Demo test program.
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

func main() {
	// case one
	// sq1 := new(Square)
	// sq1.side = 5
	// areaIntf := sq1
	// fmt.Printf("The square has area: %f\n", areaIntf.Area())

	// case two
	sq1 := new(Square)
	sq1.side = 5
	var areaIntf Shaper
	areaIntf = sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())

	// case three
	// sq1 := new(Square)
	// sq1.side = 5
	// areaIntf := Shaper(sq1)
	// fmt.Printf("The square has area: %f\n", areaIntf.Area())
}
