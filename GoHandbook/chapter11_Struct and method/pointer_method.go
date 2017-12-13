/*
@file:    pointer_method.go
@version: v1.0
@author:  aihaofeng
@date:    2017.08.23
@brief:   Pointer method test program.
          Go语言中没有类，但是可以在结构体类型或非结构体类型上定义方法。
*/

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println("")
	fmt.Println(">>>>>>>>>>")
	fmt.Println(v.X)
	fmt.Println(v.Y)
	fmt.Println("Before scaling:%+v, Abs: %v\n", v, v.Abs())

	v.Scale(5)
	fmt.Println("")
	fmt.Println(">>>>>>>>>>")
	fmt.Println(v.X)
	fmt.Println(v.Y)
	fmt.Println(v.Abs())
	fmt.Println("")
	fmt.Println("After scaling:%+v, Abs: %v\n", v, v.Abs())
}
