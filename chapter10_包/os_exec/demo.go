package main

import (
    "fmt"
    "log"
)

func main() {
    log.SetFlags(log.Flags() | log.Ldate | log.Lmicroseconds)
    log.Println(">>>>>>>>>>begin....")
    ....
    log.Println(">>>>>>>>>>end....")
}
