package main

import (
	"bufio"
	"fmt"
	"net"
	"smtpserver/commandHandler"
	"smtpserver/stateManager"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	conn.Write([]byte("220 localhost SMTP service ready\r\n"))
	var SManager *stateManager.StateManager
	SManager = stateManager.InitStateManager()

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}

		response := commandHandler.HandleCommand(SManager, message)
		conn.Write([]byte(response))
	}
}

func main() {
	ln, err := net.Listen("tcp", ":25")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer ln.Close()

	fmt.Println("SMTP server started on port 25")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
