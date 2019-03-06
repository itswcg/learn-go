package main

import (
	"fmt"
	"os"
)

func main() {
	var goos string = os.Getenv("GOOS")
	fmt.Print("The operating system is:", goos)
	path := os.Getenv("PATH")
	fmt.Print("Path is", path)
}
