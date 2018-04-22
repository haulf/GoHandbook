// @file:        os_args.go
// @author:      haulf
// @date:        2018.02.20
// @go version:  1.9
// @brief:       Os args test.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "Alice "
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}

	fmt.Println("Good Morning", who)
}
