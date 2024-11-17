package commands

import (
	"fmt"
	"io"

	"github.com/oleksandrcherevkov/cryptography/internal/console"
	"github.com/oleksandrcherevkov/cryptography/internal/crypto"
)

type AlgorithmCommand struct {
	input  io.ReadCloser
	output io.WriteCloser
}

func (s *AlgorithmCommand) Exec() (Command, error) {
	fmt.Println("Select action:")
	fmt.Println("1. Cesar algorithm")
	fmt.Println("2. Decrypt")
	response, err := console.GetString()
	if err != nil {
		return nil, err
	}
	switch response {
	case "1":
		cesar := crypto.NewCesar(4)
		return &ActionCommand{
			algorithm: cesar,
			input:     s.input,
			output:    s.output,
		}, nil
	case "2":
		frequency := crypto.NewFrequency()
		return &DecryptCommand{
			decrypter: frequency,
			input:     s.input,
			output:    s.output,
		}, nil
	}
	return ExitCommand{}, nil
}
