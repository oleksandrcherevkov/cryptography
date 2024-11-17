package crypto

import (
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

func (c Cesar) Encrypt(s *Stream) error {
	encrypt := func(char byte) byte {
		initial := strings.IndexByte(alphabet, char)
		transformed := (initial + c.delta) % len(alphabet)
		return alphabet[transformed]
	}
	return s.Pass(encrypt)
}

func (c Cesar) Decrypt(s *Stream) error {
	decrypt := func(char byte) byte {
		transformedIndex := strings.IndexByte(alphabet, char)
		initialIndex := (len(alphabet) + transformedIndex - c.delta) % len(alphabet)
		return alphabet[initialIndex]
	}
	return s.Pass(decrypt)
}
