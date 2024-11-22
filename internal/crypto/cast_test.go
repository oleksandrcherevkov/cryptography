package crypto

import (
	"testing"
)

func TestCastCorrectness(t *testing.T) {
	tc := []struct {
		name       string
		key        []byte
		plaintext  uint64
		ciphertext uint64
	}{
		{
			name:       "Should correctly encrypt plaintext with 128-bit key",
			key:        []byte{0x01, 0x23, 0x45, 0x67, 0x12, 0x34, 0x56, 0x78, 0x23, 0x45, 0x67, 0x89, 0x34, 0x56, 0x78, 0x9a},
			plaintext:  0x0123456789abcdef,
			ciphertext: 0x238b4fe5847e44b2,
		},
		{
			name:       "Should correctly encrypt plaintext with 80-bit key",
			key:        []byte{0x01, 0x23, 0x45, 0x67, 0x12, 0x34, 0x56, 0x78, 0x23, 0x45},
			plaintext:  0x0123456789abcdef,
			ciphertext: 0xeb6a711a2c02271b,
		},
		{
			name:       "Should correctly encrypt plaintext with 40-bit key",
			key:        []byte{0x01, 0x23, 0x45, 0x67, 0x12},
			plaintext:  0x0123456789abcdef,
			ciphertext: 0x7ac816d16e9b302e,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			cast, err := NewCast(tt.key)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			encrypted := cast.cast(tt.plaintext)
			if encrypted != tt.ciphertext {
				t.Fatalf("Expected %x, got %x", tt.ciphertext, encrypted)
			}
		})
	}
}

func TestEncryptDecrypt(t *testing.T) {
	tc := []struct {
		name      string
		key       []byte
		plaintext uint64
	}{
		{
			name:      "Should correctly encrypt plaintext with 128-bit key",
			key:       []byte{0x01, 0x23, 0x45, 0x67, 0x12, 0x34, 0x56, 0x78, 0x23, 0x45, 0x67, 0x89, 0x34, 0x56, 0x78, 0x9a},
			plaintext: 0x0123456789abcdef,
		},
		{
			name:      "Should correctly encrypt plaintext with 80-bit key",
			key:       []byte{0x01, 0x23, 0x45, 0x67, 0x12, 0x34, 0x56, 0x78, 0x23, 0x45},
			plaintext: 0x0123456789abcdef,
		},
		{
			name:      "Should correctly encrypt plaintext with 40-bit key",
			key:       []byte{0x01, 0x23, 0x45, 0x67, 0x12},
			plaintext: 0x0123456789abcdef,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			cast, err := NewCast(tt.key)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			encrypted := cast.cast(tt.plaintext)
			decrypted := cast.recast(encrypted)
			if decrypted != tt.plaintext {
				t.Fatalf("Expected %x, got %x", tt.plaintext, encrypted)
			}
		})
	}
}

func TestExtractUint64(t *testing.T) {
	tc := []struct {
		name     string
		key      key
		keyFirst int
		expected uint32
	}{
		{
			name:     "Should correctly encrypt plaintext with 128-bit key",
			key:      [16]byte{0x01, 0x23, 0x45, 0x67, 0x12, 0x34, 0x56, 0x78, 0x23, 0x45, 0x67, 0x89, 0x34, 0x56, 0x78, 0x9a},
			keyFirst: 0x0,
			expected: 0x01234567,
		},
		{
			name:     "Should correctly encrypt plaintext with 80-bit key",
			key:      [16]byte{0x01, 0x23, 0x45, 0x67, 0x12, 0x34, 0x56, 0x78, 0x23, 0x45, 0x67, 0x89, 0x34, 0x56, 0x78, 0x9a},
			keyFirst: 0x4,
			expected: 0x12345678,
		},
		{
			name:     "Should correctly encrypt plaintext with 40-bit key",
			key:      [16]byte{0x01, 0x23, 0x45, 0x67, 0x12, 0x34, 0x56, 0x78, 0x23, 0x45, 0x67, 0x89, 0x34, 0x56, 0x78, 0x9a},
			keyFirst: 0x8,
			expected: 0x23456789,
		},
		{
			name:     "Should correctly encrypt plaintext with 40-bit key",
			key:      [16]byte{0x01, 0x23, 0x45, 0x67, 0x12, 0x34, 0x56, 0x78, 0x23, 0x45, 0x67, 0x89, 0x34, 0x56, 0x78, 0x9a},
			keyFirst: 0xc,
			expected: 0x3456789a,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			extracted := tt.key.extractUint32(tt.keyFirst)
			if extracted != tt.expected {
				t.Fatalf("Expected %b, got %b", tt.expected, extracted)
			}
		})
	}
}

func TestRotation(t *testing.T) {
	tc := []struct {
		name     string
		number   uint32
		expected byte
	}{
		{
			name:     "Should correctly encrypt plaintext with 128-bit key",
			number:   0xFFFFFFFF,
			expected: 0b00011111,
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			rotation := rotation(tt.number)
			if rotation != tt.expected {
				t.Fatalf("Expected %b, got %b", tt.expected, rotation)
			}
		})
	}
}
