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
	fmt.Println("3. Gamma")
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
	case "3":
		fmt.Print("Input A -> ")
		a, err := console.GetUint()
		if err != nil {
			return nil, err
		}
		fmt.Print("Input C -> ")
		c, err := console.GetUint()
		if err != nil {
			return nil, err
		}
		fmt.Print("Input T(0) (seed random value) -> ")
		t, err := console.GetUint()
		if err != nil {
			return nil, err
		}
		gamma := crypto.NewGamma(a, c, t)
		return &ActionCommand{
			algorithm: gamma,
			input:     s.input,
			output:    s.output,
		}, nil
	}
	return ExitCommand{}, nil
}
