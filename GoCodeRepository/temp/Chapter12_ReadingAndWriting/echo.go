// file:   ehco.go
// author: haulf
// date:   20180527
// brief:  Simulate the Unix echo-utility.

package main

import (
	"flag"
	"fmt"
	"os"
)

var NewLine = flag.Bool("n", false, "print on newline")

const (
	Space   = " "
	Newline = "\n"
)

func main() {
	flag.PrintDefaults()

	flag.Parse() // Scan the argument list(or list of constants) and sets up flags.

	var s string = " "
	fmt.Println("The number of arguments is:", flag.NArg())
	for i := 0; i < flag.NArg(); i++ { // flag.NArg() is the number of arguments.
		if i > 0 {
			s += Space
		}
		s += flag.Arg(i)

		if *NewLine {
			s += Newline
		}
	}

	// if *NewLine { // -n is parsed, flag becomes true
	// s += Newline
	// }

	os.Stdout.WriteString(s)
}
