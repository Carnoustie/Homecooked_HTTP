package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("\n\nHTTP Server is running")

	listen, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("\n\nListen failed with error: ", err)
	}

	fmt.Println("\n\nListening on port 8000\n\n")
}
