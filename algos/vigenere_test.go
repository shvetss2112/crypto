package algos

import (
	"bytes"
	"testing"
)

func TestVigenereEncrypt(t *testing.T) {
	tests := []struct {
		name      string
		plaintext []byte
		key       string
		expected  []byte
	}{
		{
			name:      "simple key",
			plaintext: []byte("hello"),
			key:       "key",
			expected:  []byte{104 + 107, 101 + 101, 108 + 121, 108 + 107, 111 + 101},
		},
		{
			name:      "single byte",
			plaintext: []byte("a"),
			key:       "x",
			expected:  []byte{97 + 120},
		},
		{
			name:      "key repeats",
			plaintext: []byte("abcdef"),
			key:       "ab",
			expected:  []byte{97 + 97, 98 + 98, 99 + 97, 100 + 98, 101 + 97, 102 + 98},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encFunc := Vigenere(true, tt.key)
			output := make([]byte, len(tt.plaintext))
			encFunc(tt.plaintext, output)

			if !bytes.Equal(output, tt.expected) {
				t.Errorf("Vigenere encryption failed. Got %v, expected %v", output, tt.expected)
			}
		})
	}
}

func TestVigenereDecrypt(t *testing.T) {
	tests := []struct {
		name       string
		ciphertext []byte
		key        string
		expected   []byte
	}{
		{
			name:       "simple key",
			ciphertext: []byte{104 + 107, 101 + 101, 108 + 121, 108 + 107, 111 + 101},
			key:        "key",
			expected:   []byte("hello"),
		},
		{
			name:       "single byte",
			ciphertext: []byte{97 + 120},
			key:        "x",
			expected:   []byte("a"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decFunc := Vigenere(false, tt.key)
			output := make([]byte, len(tt.ciphertext))
			decFunc(tt.ciphertext, output)

			if !bytes.Equal(output, tt.expected) {
				t.Errorf("Vigenere decryption failed. Got %v, expected %v", output, tt.expected)
			}
		})
	}
}

func TestVigenereRoundTrip(t *testing.T) {
	plaintext := []byte("The quick brown fox jumps over the lazy dog")
	keys := []string{"key", "secret", "a", "vigenerekey123"}

	for _, key := range keys {
		t.Run("roundtrip_key_"+key, func(t *testing.T) {
			encFunc := Vigenere(true, key)
			decFunc := Vigenere(false, key)

			encrypted := make([]byte, len(plaintext))
			decrypted := make([]byte, len(plaintext))

			encFunc(plaintext, encrypted)
			decFunc(encrypted, decrypted)

			if !bytes.Equal(plaintext, decrypted) {
				t.Errorf("Round trip failed. Original: %v, Result: %v", plaintext, decrypted)
			}
		})
	}
}

