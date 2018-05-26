// @file:        ReadInput1.go
// @version:     1.0
// @date:        2017.12.11
// @brief:       Standard input and output test.

package main

import (
	"fmt"
)

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	inputString            = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

func main() {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	// fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Your name is: %s %s!\n", firstName, lastName)

	fmt.Sscanf(inputString, format, &f, &i, &s)
	fmt.Println("From the inputString, we read: ", f, i, s)
}
