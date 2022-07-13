package commands

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/client-file-upload/models"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func GetInput() (string, error) {
	fmt.Print("-> ")

	str, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	str = strings.Replace(str, "\n", "", 1)

	return str, nil
}

func RunCommand(data string, conn net.Conn) {

	if strings.HasPrefix(data, "send") {
		commands := strings.Split(data, " ")

		fileName := commands[1]

		SendFile(fileName, conn)
		return
	}

	gob.NewEncoder(conn).Encode(&models.Messages{Message: data})
}

func ReadCommand(data models.Messages) {
	if strings.HasPrefix(data.Message, "send") {
		command := strings.Split(data.Message, " ")
		fileName := command[1]

		ReceiveFile(fileName, data.Args)
		return
	}
}

func SendFile(fileName string, conn net.Conn) {
	fmt.Printf("Sending file %s...\n", fileName)

	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := models.Messages{Message: "send " + fileName, Args: file}

	err = gob.NewEncoder(conn).Encode(&data)
	if err != nil {
		fmt.Println(err)
	}
}

func ReceiveFile(fileName string, fileBuffer []byte) {
	fmt.Println("Receiving file...")
	file, err := os.Create(fileName)

	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	_, err = file.Write(fileBuffer)

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("File received.")

}
