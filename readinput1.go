package main

//import (
//	"bufio"
//	"fmt"
//	"os"
//)

//var (
//	firstName, lastName, s string
//	i                      int
//	f                      float32
//	input                  = "56.12/5252/Go"
//	format                 = "%f/%d/%s"
//)
//
//func main() {
//	fmt.Println("Please enter your full name:")
//	fmt.Scanln(&firstName, &lastName)
//	fmt.Printf("Hi %s %s\n", firstName, lastName)
//	fmt.Sscanf(input, format, &f, &i, &s)
//	fmt.Println("From the string we read:", f, i, s)
//}

//var inputReader *bufio.Reader
//var input string
//var err error
//
//func main() {
//	inputReader = bufio.NewReader(os.Stdin)
//	fmt.Println("Please enter some input:")
//	input, err = inputReader.ReadString('\n')
//	if err == nil {
//		fmt.Printf("The input was: %s\n", input)
//	}
//}

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "Alice"
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Good", who)
}
