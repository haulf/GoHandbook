package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(".apk{1}")
	fmt.Println(re.MatchString("gopher.apk"))
	fmt.Println(re.MatchString("gophergopherdf"))
	fmt.Println(re.MatchString("gophergophergopherddd"))
}
