package commands

import (
	"errors"
	"fmt"
	"io"
	"strconv"

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
	fmt.Println("4. CAST-128")
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
	case "4":
		fmt.Print("Input key hex -> ")
		key, err := getCastKey()
		if err != nil {
			return nil, err
		}
		cast := crypto.NewCast(key)
		return &ActionCommand{
			algorithm: cast,
			input:     s.input,
			output:    s.output,
		}, nil
	}
	return ExitCommand{}, nil
}

func getCastKey() ([16]byte, error) {
	var key [16]byte
	keyString, err := console.GetString()
	if err != nil {
		return key, err
	}
	if len(keyString) != 32 {
		return key, errors.New("key not right size of 32 hex values")
	}
	for i := 0; i < 32; i += 2 {
		byteString := keyString[i : i+2]
		number, err := strconv.ParseUint(byteString, 16, 8)
		if err != nil {
			return key, err
		}
		key[i/2] = byte(number)
	}
	return key, nil
}
