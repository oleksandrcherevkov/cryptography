package commands

import (
	"fmt"
	"io"

	"github.com/oleksandrcherevkov/cryptography/internal/console"
	"github.com/oleksandrcherevkov/cryptography/internal/crypto"
)

type ActionCommand struct {
	algorithm crypto.Algorithm
	input     io.ReadCloser
	output    io.WriteCloser
}

func (s *ActionCommand) Exec() (Command, error) {
	fmt.Println("Select action:")
	fmt.Println("1. Encrypt")
	fmt.Println("2. Decrypt")
	response, err := console.GetString()
	if err != nil {
		return nil, err
	}
	switch response {
	case "1":
		err = s.algorithm.Encrypt(s.input, s.output)
	case "2":
		err = s.algorithm.Decrypt(s.input, s.output)
	}
	if err != nil {
		return nil, err
	}
	s.input.Close()
	s.output.Close()
	return ExitCommand{}, nil
}
