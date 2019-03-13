package main

import "fmt"

type TwoInts struct {
	a int
	b int
}

//func main() {
//	two1 := new(TwoInts)
//	two1.a = 12
//
//	fmt.Printf("The sum is: %d\n", two1.AddThem())
//	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))
//
//	two2 := TwoInts{3, 4}
//	fmt.Printf("The sum is: %d\n", two2.AddThem())
//
//}
//
//func (tn *TwoInts) AddThem() int {
//	return tn.a + tn.b
//}
//
//func (tn *TwoInts) AddToParam(param int) int {
//	return tn.a + tn.b + param
//}

//type IntVector []int
//
//func (v IntVector) Sum() (s int) {
//	for _, x := range v {
//		s += x
//	}
//	return
//}
//
//func main() {
//	fmt.Println(IntVector{1, 2, 3}.Sum())
//}

type B struct {
	thing int
}

func (b *B) change() {
	b.thing = 1
}

func (b B) write() string {
	return fmt.Sprint(b)
}

func main() {
	var b1 B
	b1.change()
	fmt.Println(b1.write())

	b2 := new(B)
	b2.change()
	fmt.Println(b2.write())
}
