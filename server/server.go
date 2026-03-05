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

func Message(message *gen.ClientMessage) (*gen.ServerMessage, error) {
	s := ChatServer{}
	var ctx context.Context
	serverMessage, err := s.Chat(ctx, message)
	if err != nil {
		return nil, err
	}
	return serverMessage, nil
}
