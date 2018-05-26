/*
@file:    normal_function.go
@version: v1.0
@author:  aihaofeng
@date:    2017.08.23
@brief:   Normal function test program.
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

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println("")
	fmt.Println(v.X)
	fmt.Println(v.Y)
	fmt.Println(Abs(v))

	Scale(&v, 10) // 函数Scale()的第一个参数是一个指针。
	fmt.Println("")
	fmt.Println(v.X)
	fmt.Println(v.Y)
	fmt.Println(Abs(v))
	fmt.Println("")
}
