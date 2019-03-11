package main

import (
	"fmt"
	"strings"
)

//type struct1 struct {
//	i1  int
//	f1  float32
//	str string
//}
//
//func main() {
//	ms := new(struct1)
//	ms.i1 = 10
//	ms.f1 = 15.5
//	ms.str = "wcg"
//
//	fmt.Println(ms)
//}

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	var pers1 Person
	pers1.firstName = "its"
	pers1.lastName = "wcg"
	upPerson(&pers1)
	fmt.Printf("The name of person is %s %s\n", pers1.firstName, pers1.lastName)

	pers2 := new(Person)
	pers2.firstName = "its"
	pers2.lastName = "wcg"
	(*pers2).lastName = "wcg"
	upPerson(pers2)

	pers3 := &Person{"its", "wcg"}
	upPerson(pers3)
}
