package main

import (
	"fmt"
	"time"
)

//func main() {
//	result := 0
//	for i := 0; i <= 10; i++ {
//		result = fibonacc(i)
//		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
//	}
//}
//
//func fibonacc(n int) (res int) {
//	if n <= 1 {
//		res = 1
//	} else {
//		res = fibonacc(n-1) + fibonacc(n-2)
//	}
//	return
//}
//

const LIM = 41

var fibs [LIM]uint64

func main() {
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

func fibonacci(n int) (res uint64) {
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return
}

//func main() {
//	callback(1, Add)
//}
//
//func Add(a, b int) {
//	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
//}
//
//func callback(y int, f func(int, int)) {
//	f(y, 2)
//}

//func main() {
//	p2 := Add2()
//	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(2))
//
//	TwoAdder := Adder(2)
//	fmt.Printf("The result is :%v\n", TwoAdder(2))
//}
//
//func Add2() func(b int) int {
//	return func(b int) int {
//		return b + 2
//	}
//}
//
//func Adder(a int) func(b int) int {
//	return func(b int) int {
//		return a + b
//	}
//}

//func MakeAddSuffix(suffix string) func(string) string {
//	return func(name string) string {
//		if !string.HasSuffix(name, suffix) {
//			return name + suffix
//		}
//		return name
//	}
//}
//
//addBmp := MakeAddSuffix(".bmp")
//addJpeg := MakeAddSuffix(".jpeg")
//
//addBmp("file")
//addJpeg("file")
