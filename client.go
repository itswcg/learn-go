package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	clientName, _ := inputReader.ReadString('\n')
	trimmendClient := strings.Trim(clientName, "\n")

	for {
		fmt.Println("What to send to the server? Type Q to quit")
		input, _ := inputReader.ReadString('\n')
		trimmendInput := strings.Trim(input, "\n")
		if trimmendInput == "Q" {
			return
		}

		_, err = conn.Write([]byte(trimmendClient + "says:" + trimmendInput))
	}
}
