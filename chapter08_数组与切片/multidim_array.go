/*
@file:    multidim_array.go
@version: v1.0
@author:  aihaofeng
@date:    2017.11.12
@brief:   Array test program.
*/

package main

const (
	WIDTH  = 1920
	HEIGHT = 1080
)

type pixel int

var screen [WIDTH][HEIGHT]pixel

func main() {

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = 0
		}
	}
}
