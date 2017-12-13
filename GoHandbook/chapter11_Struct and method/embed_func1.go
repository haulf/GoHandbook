// @file:        embed_func1.go
// @version:     1.0
// @author:      aihaofeng
// @date:        2017.12.10
// @go version:  1.9
// @brief:       Function test.

package main

import (
	"fmt"
)

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log  *Log
}

func main() {
	c := new(Customer)
	c.Name = "Barak Obama"
	c.log = new(Log)
	c.log.msg = "1 - Yes we can!"

	// 上面的代码可以用下面更加简洁的方式代替
	//c = &Customer{"Barak Obama", &Log{"1 - Yes wecan!"}}

	c.Log().Add("2 - After me the world will be a better place!")

	fmt.Println(c.Log().String())
	//fmt.Println(c.Log())
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) Log() *Log {
	return c.log
}
