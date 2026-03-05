package client

import (
	"bufio"
	"grpcchat/gen/gen"
	"os"
)

func Message() *gen.ClientMessage {
	scanner := bufio.NewScanner(os.Stdin)
	var message string
	for scanner.Scan() {
		message = scanner.Text()
	}
	println(message)
	return nil
}
