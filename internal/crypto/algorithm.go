package crypto

import "io"

type Algorithm interface {
	Encrypt(io.Reader, io.Writer) error
	Decrypt(io.Reader, io.Writer) error
}
