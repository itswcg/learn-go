package main

import "fmt"

func Multiply(a, b int, reply *int) {
	*reply = a * b
}

func main() {
	n := 0
	relpay := &n
	Multiply(10, 5, relpay)
	fmt.Println("Multipy", *relpay)
}
