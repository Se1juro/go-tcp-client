package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/client-file-upload/commands"
	"github.com/client-file-upload/models"
)

func client(network string) {
	connection, error := net.Dial("tcp", network)
	if error != nil {
		fmt.Println(error)
		return
	}
	log.Printf("Connected to server %s", network)

	go readMessages(connection)

	writeMessage(connection)
}

func readMessages(conn net.Conn) {
	for {
		var data models.Messages
		err := gob.NewDecoder(conn).Decode(&data)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%+v\n", data.Message)
		commands.ReadCommand(data)
	}
}

func writeMessage(conn net.Conn) {
	for {
		input, _ := commands.GetInput()
		fmt.Println(input)
		if input == "exit" {
			break
		}
		commands.RunCommand(input, conn)
	}
}

func main() {
	var startServer string
	flag.StringVar(&startServer, "start", "", "Connect to TCP server")
	flag.Parse()
	client(":" + startServer)
}
