package main

import (
	"grpcchat/client"
	"grpcchat/server"
	"log"
)

func main() {
	clientMessage := client.Message()
	serverMessage, err := server.Message(clientMessage)
	if err != nil {
		log.Fatal(err)
	}
	println(serverMessage.Text)
}
