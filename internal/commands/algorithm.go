package commands

import (
	"io"

	"github.com/oleksandrcherevkov/cryptography/internal/crypto"
)

type AlgorithmCommand struct {
	input  io.ReadCloser
	output io.WriteCloser
}

func (s *AlgorithmCommand) Exec() (Command, error) {
	cesar := crypto.NewCesar(4)
	return &ActionCommand{
		algorithm: cesar,
		input:     s.input,
		output:    s.output,
	}, nil
}
