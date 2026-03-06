package client

import (
	"bufio"
	"grpcchat/gen/gen"
	"os"
)

func Message() *gen.ClientMessage {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	for scanner.Scan() {
		input = scanner.Text()
		break
	}
	var clientMessage gen.ClientMessage
	if input != "" {
		clientMessage = gen.ClientMessage{
			Username: "test",
			Text:     input,
		}
	}
	return &clientMessage
}
