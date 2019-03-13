package main

import (
	"fmt"
	"strconv"
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
	c.Name = "wcg"
	c.log = new(Log)
	c.log.msg = "I am wcg"

	//c = &Customer{"wcg", &Log{"I am wcg"}}
	//c.Log().Add("md")
	//fmt.Println(c.log())
	fmt.Println(c)
}

type Integer int

func (i *Integer) String() string {
	return strconv.Itoa(i)
}
