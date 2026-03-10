package main

import (
	"grpcchat/gen/gen"
	"log"

	"google.golang.org/grpc"
)

var client gen.ChatServiceClient

func main() {
	conn, err := grpc.NewClient("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client = gen.NewChatServiceClient(conn)
}
