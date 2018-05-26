// file:    5.5.2.6_Random.go
// author:  haulf
// date:    2017.08.27
// brief:   Random test program.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(">>> 10 random numbers:")
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d \n", a)
	}
	fmt.Println("")

	fmt.Println(">>> 5 random numbers:")
	for i := 0; i < 5; i++ {
		r := rand.Intn(8)
		fmt.Printf("%d\n ", r)
	}
	fmt.Println("")

	fmt.Println(">>> 10 random numbers:")
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f\n ", 100*rand.Float32())
	}
	fmt.Println("")
}
