package algos

import (
	"bytes"
	"testing"
)

func TestFenceEncrypt(t *testing.T) {
	tests := []struct {
		name      string
		plaintext []byte
		key       string
	}{
		{
			name:      "simple text with key 2",
			plaintext: []byte("hello"),
			key:       "2",
		},
		{
			name:      "simple text with key 3",
			plaintext: []byte("hello"),
			key:       "3",
		},
		{
			name:      "longer text with key 2",
			plaintext: []byte("thequickbrownfox"),
			key:       "2",
		},
		{
			name:      "single byte",
			plaintext: []byte("a"),
			key:       "2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encFunc := Fence(true, tt.key)
			output := make([]byte, len(tt.plaintext))
			encFunc(tt.plaintext, output)

			if len(output) != len(tt.plaintext) {
				t.Errorf("Output length mismatch. Got %d, expected %d", len(output), len(tt.plaintext))
			}
		})
	}
}

func TestFenceRoundTrip(t *testing.T) {
	tests := []struct {
		name      string
		plaintext []byte
		key       string
	}{
		{
			name:      "key 2",
			plaintext: []byte("thequickbrownfox"),
			key:       "2",
		},
		{
			name:      "key 3",
			plaintext: []byte("thequickbrownfoxjumps"),
			key:       "3",
		},
		{
			name:      "key 4",
			plaintext: []byte("abcdefghijklmnop"),
			key:       "4",
		},
		{
			name:      "key 2 short text",
			plaintext: []byte("abc"),
			key:       "2",
		},
		{
			name:      "key 3 short text",
			plaintext: []byte("abcde"),
			key:       "3",
		},
		{
			name:      "single character",
			plaintext: []byte("a"),
			key:       "2",
		},
		{
			name:      "two characters",
			plaintext: []byte("ab"),
			key:       "2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encFunc := Fence(true, tt.key)
			decFunc := Fence(false, tt.key)

			encrypted := make([]byte, len(tt.plaintext))
			decrypted := make([]byte, len(tt.plaintext))

			encFunc(tt.plaintext, encrypted)
			decFunc(encrypted, decrypted)

			if !bytes.Equal(tt.plaintext, decrypted) {
				t.Errorf("Round trip failed. Original: %v, Result: %v", tt.plaintext, decrypted)
			}
		})
	}
}


func TestFenceByteBoundaries(t *testing.T) {
	plaintext := []byte{0, 127, 255, 1, 128, 64, 32}
	key := "2"

	encFunc := Fence(true, key)
	decFunc := Fence(false, key)

	encrypted := make([]byte, len(plaintext))
	decrypted := make([]byte, len(plaintext))

	encFunc(plaintext, encrypted)
	decFunc(encrypted, decrypted)

	if !bytes.Equal(plaintext, decrypted) {
		t.Errorf("Byte boundary test failed. Original: %v, Result: %v", plaintext, decrypted)
	}
}
