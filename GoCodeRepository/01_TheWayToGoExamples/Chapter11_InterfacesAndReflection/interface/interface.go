// file:    interface.go
// author:  haulf
// date:    2017.09.28
// brief:   Demo test program.

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
	// square := new(Square)
	// square.side = 5
	// areaInterface := square
	// fmt.Printf("The square has area: %f\n", areaInterface.Area())

	// case two
	// square := new(Square)
	// square.side = 5
	// var areaInterface Shaper
	// areaInterface = square
	// fmt.Printf("The square has area: %f\n", areaInterface.Area())

	// case three
	square := new(Square)
	square.side = 5
	areaInterface := Shaper(square)
	fmt.Printf("The square has area: %f\n", areaInterface.Area())
}
