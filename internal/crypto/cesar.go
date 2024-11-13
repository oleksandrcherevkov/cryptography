package crypto

import (
	"errors"
	"io"
	"strings"
)

var alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKMLNOPQRSTUVWXYZ "

type Cesar struct {
	delta int
}

func NewCesar(delta int) Cesar {
	return Cesar{
		delta,
	}
}

func (c Cesar) Encrypt(input io.Reader, output io.Writer) error {
	encrypt := func(char byte) byte {
		initial := strings.IndexByte(alphabet, char)
		transformed := (initial + c.delta) % len(alphabet)
		return alphabet[transformed]
	}
	return iterate(input, output, encrypt)
}

func (c Cesar) Decrypt(input io.Reader, output io.Writer) error {
	decrypt := func(char byte) byte {
		transformedIndex := strings.IndexByte(alphabet, char)
		initialIndex := (len(alphabet) + transformedIndex - c.delta) % len(alphabet)
		return alphabet[initialIndex]
	}
	return iterate(input, output, decrypt)
}

func iterate(input io.Reader, output io.Writer, transform func(byte) byte) error {
	char := make([]byte, 1)
	for {
		n, err := input.Read(char)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if n != 1 {
			return errors.New("no read from input")
		}
		if char[0] == '\n' {
			break
		}
		char[0] = transform(char[0])
		n, err = output.Write(char)
		if err != nil {
			return err
		}
		if n != 1 {
			return errors.New("no write to output")
		}
	}
	char[0] = '\n'
	n, err := output.Write(char)
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("no write to output")
	}
	return nil

}
