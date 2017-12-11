/*
@file:    5.5.2.1_type_mixing.go
@version: v1.0
@author:  aihaofeng
@date:    2017.08.27
@brief:   Type mixing test program.
*/

package main

func main() {
	var a int
	var b int32
	a = 15
	b = a + a // 编译错误
	b = b + 5 // 因为5是常量,所以可以通过编译
}
