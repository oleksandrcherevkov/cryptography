package crypto

import (
	"bytes"
	"errors"
	"io"
)

type Stream struct {
	reader       io.Reader
	readerBuffer bytes.Buffer
	writer       io.Writer
}

func NewStream(reader io.Reader, writer io.Writer) *Stream {
	stream := &Stream{}
	stream.reader = io.TeeReader(reader, &stream.readerBuffer)
	stream.writer = writer
	return stream
}

func (s *Stream) Read(b []byte) (int, error) {
	return s.reader.Read(b)
}

func (s *Stream) Write(b []byte) (int, error) {
	return s.writer.Write(b)
}

func (s *Stream) Pass(transform func(byte) byte) error {
	char := make([]byte, 1)
	for {
		n, err := s.Read(char)
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
		n, err = s.Write(char)
		if err != nil {
			return err
		}
		if n != 1 {
			return errors.New("no write to output")
		}
	}
	s.refreshReader()
	char[0] = '\n'
	n, err := s.Write(char)
	if err != nil {
		return err
	}
	if n != 1 {
		return errors.New("no write to output")
	}
	return nil

}

func (s *Stream) Scan(f func(byte) error) error {
	char := make([]byte, 1)
	for {
		n, err := s.Read(char)
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
		err = f(char[0])
		if err != nil {
			return err
		}
	}
	s.refreshReader()
	return nil
}

func (s *Stream) refreshReader() {
	oldBuffer := s.readerBuffer
	s.readerBuffer = bytes.Buffer{}
	s.reader = io.TeeReader(&oldBuffer, &s.readerBuffer)
}
