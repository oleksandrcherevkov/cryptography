package crypto

type Encrypter interface {
	Encrypt(*Stream) error
}

type Decrypter interface {
	Decrypt(*Stream) error
}

type Algorithm interface {
	Encrypter
	Decrypter
}
