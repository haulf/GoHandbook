// @file:        PackageRegexp.go
// @author:      haulf
// @date:        2017.11.20
// @brief:       Regexp test.

package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	originalString := "John: 2578.34 William: 4567.23 Steve: 5632.18"

    fmt.Println("Original string is:", originalString)

    fmt.Println("Start to match")
	patternString := "[0-9]+.[0-9]+"
	if ok, _ := regexp.Match(patternString, []byte(originalString)); ok {
		fmt.Println("Match Found!")
	}

	regexpObject, _ := regexp.Compile(patternString)

	str := regexpObject.ReplaceAllString(originalString, "##.#")
	fmt.Println(str)

	paramFunc := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}
	strFunc := regexpObject.ReplaceAllStringFunc(originalString, paramFunc)
	fmt.Println(strFunc)
}
