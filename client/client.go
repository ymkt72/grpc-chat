package client

import (
	"bufio"
	"os"
)

func Message() {
	scanner := bufio.NewScanner(os.Stdin)
	var message string
	for scanner.Scan() {
		message = scanner.Text()
	}
	println(message)
}
