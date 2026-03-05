package server

import (
	"context"
	"grpcchat/gen/gen"
)

type ChatServer struct{}

func (c *ChatServer) Chat(context.Context, *gen.ClientMessage) (*gen.ServerMessage, error) {
	serverMessage := gen.ServerMessage{
		From:      "userA",
		Text:      "you are dumb!",
		Timestamp: nil,
	}
	return &serverMessage, nil
}

func Message() (*gen.ServerMessage, error) {
	s := ChatServer{}
	var ctx context.Context
	clientMessage := gen.ClientMessage{Username: "test", Text: "hello"}
	serverMessage, err := s.Chat(ctx, &clientMessage)
	if err != nil {
		return nil, err
	}
	return serverMessage, nil
}
