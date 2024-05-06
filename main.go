package main

import (
	"fmt"
	"os"

	"01.alem.school/git/dbaitako/net-cat/internal"
)

func main() {
	var port string
	switch {
	case len(os.Args) == 1:
		port = "8989"
	case len(os.Args) == 2:
		port = os.Args[1]
	default:
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	internal.StartServer(port)
}
