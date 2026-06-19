package algos

import (
	"bytes"
	"testing"
)

func TestAutokeyEncrypt(t *testing.T) {
	tests := []struct {
		name      string
		plaintext []byte
		key       string
	}{
		{
			name:      "simple plaintext",
			plaintext: []byte("hello"),
			key:       "key",
		},
		{
			name:      "single byte",
			plaintext: []byte("a"),
			key:       "x",
		},
		{
			name:      "key longer than plaintext",
			plaintext: []byte("hi"),
			key:       "verylongkey",
		},
		{
			name:      "plaintext longer than key",
			plaintext: []byte("verylongplaintext"),
			key:       "k",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encFunc := Autokey(true, tt.key)
			output := make([]byte, len(tt.plaintext))
			encFunc(tt.plaintext, output)

			// Just verify it produces output of the correct length
			if len(output) != len(tt.plaintext) {
				t.Errorf("Output length mismatch. Got %d, expected %d", len(output), len(tt.plaintext))
			}
		})
	}
}

func TestAutokeyRoundTrip(t *testing.T) {
	tests := []struct {
		name      string
		plaintext []byte
		key       string
	}{
		{
			name:      "simple roundtrip",
			plaintext: []byte("hello"),
			key:       "key",
		},
		{
			name:      "longer text",
			plaintext: []byte("The quick brown fox jumps over the lazy dog"),
			key:       "secret",
		},
		{
			name:      "single byte",
			plaintext: []byte("a"),
			key:       "x",
		},
		{
			name:      "plaintext longer than key",
			plaintext: []byte("supercalifragilisticexpialidocious"),
			key:       "a",
		},
		{
			name:      "key longer than plaintext",
			plaintext: []byte("hi"),
			key:       "verylongkeystring",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encFunc := Autokey(true, tt.key)
			decFunc := Autokey(false, tt.key)

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


func TestAutokeyByteBoundaries(t *testing.T) {
	plaintext := []byte{0, 127, 255, 1, 128}
	key := "test"

	encFunc := Autokey(true, key)
	decFunc := Autokey(false, key)

	encrypted := make([]byte, len(plaintext))
	decrypted := make([]byte, len(plaintext))

	encFunc(plaintext, encrypted)
	decFunc(encrypted, decrypted)

	if !bytes.Equal(plaintext, decrypted) {
		t.Errorf("Byte boundary test failed. Original: %v, Result: %v", plaintext, decrypted)
	}
}
