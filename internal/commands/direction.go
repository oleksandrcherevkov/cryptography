package commands

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/oleksandrcherevkov/cryptography/internal/console"
)

type DirectionCommand struct{}

func (c DirectionCommand) Exec() (Command, error) {
	input, err := selectInput()
	if err != nil {
		return nil, err
	}

	output, err := selectOutput()
	if err != nil {
		return nil, err
	}

	return &AlgorithmCommand{
		input:  input,
		output: output,
	}, nil
}

func selectInput() (io.ReadCloser, error) {
	fmt.Println("Select text source:")
	fmt.Println("1. Console")
	fmt.Println("2. File")
	response, err := console.GetString()
	if err != nil {
		return nil, err
	}

	switch response {
	case "1":
		fmt.Println("Input text:")
		response, err = console.GetString()
		if err != nil {
			return nil, err
		}
		return &readEmptyCloser{strings.NewReader(response)}, nil
	case "2":
		return console.GetFile()
	default:
		fmt.Println(response, " is not right response")
	}

	return nil, errors.New("empty input")
}

func selectOutput() (io.WriteCloser, error) {
	fmt.Println("Select text destination:")
	fmt.Println("1. Console")
	fmt.Println("2. File")
	response, err := console.GetString()
	if err != nil {
		return nil, err
	}

	switch response {
	case "1":
		return &writeEmptyCloser{os.Stdout}, nil
	case "2":
		return console.GetFile()
	}

	return nil, nil
}

type readEmptyCloser struct {
	reader io.Reader
}

func (r *readEmptyCloser) Read(b []byte) (int, error) {
	return r.reader.Read(b)
}

func (r *readEmptyCloser) Close() error {
	return nil
}

type writeEmptyCloser struct {
	writer io.Writer
}

func (r *writeEmptyCloser) Write(b []byte) (int, error) {
	return r.writer.Write(b)
}

func (r *writeEmptyCloser) Close() error {
	return nil
}
