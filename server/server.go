package main

import (
	"context"
	"grpcchat/gen/gen"
	"log"
	"net"

	"google.golang.org/grpc"
)

type ChatServer struct {
	gen.UnimplementedChatServiceServer
}

func (c ChatServer) Chat(context.Context, *gen.ClientMessage) (*gen.ServerMessage, error) {
	serverMessage := gen.ServerMessage{
		From:      "userA",
		Text:      "you are dumb!",
		Timestamp: nil,
	}
	return &serverMessage, nil
}

func newServer() ChatServer {
	return ChatServer{}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	gen.RegisterChatServiceServer(grpcServer, newServer())

	grpcServer.Serve(lis)
}
