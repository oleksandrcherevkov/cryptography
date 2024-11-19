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
	stream := crypto.NewStream(s.input, s.output, false)
	switch response {
	case "1":
		err = s.algorithm.Encrypt(stream)
	case "2":
		err = s.algorithm.Decrypt(stream)
	}
	if err != nil {
		return nil, err
	}
	s.input.Close()
	s.output.Close()
	return ExitCommand{}, nil
}

type DecryptCommand struct {
	decrypter crypto.Decrypter
	input     io.ReadCloser
	output    io.WriteCloser
}

func (s *DecryptCommand) Exec() (Command, error) {
	stream := crypto.NewStream(s.input, s.output, true)
	err := s.decrypter.Decrypt(stream)
	if err != nil {
		return nil, err
	}
	s.input.Close()
	s.output.Close()
	return ExitCommand{}, nil
}
