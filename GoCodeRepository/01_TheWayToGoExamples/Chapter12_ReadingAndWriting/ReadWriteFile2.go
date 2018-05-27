// file:        ReadWriteFile2.go
// author:      haulf
// date:        2018.02.20
// brief:       If the data columns are separated by a space, you can use the FScan-function
//              series from the "fmt" package. This is applied in  the following program,
//              which reads in data form 3 columns into the variables v1, v2 and v3; they are
//              then appended to column slices.

package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("products2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var col1, col2, col3 []string

	for {
		var v1, v2, v3 string
		// scans until newline
		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
		if err != nil {
			break
		}

		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}

	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}
