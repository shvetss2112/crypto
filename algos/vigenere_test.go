package algos

import (
	"bytes"
	"testing"
)

func TestVigenereRoundTrip(t *testing.T) {
	plaintext := []byte("The quick brown fox")
	key := "secret"

	encFunc := Vigenere(true, key)
	decFunc := Vigenere(false, key)

	encrypted := make([]byte, len(plaintext))
	decrypted := make([]byte, len(plaintext))

	encFunc(plaintext, encrypted)
	decFunc(encrypted, decrypted)

	if !bytes.Equal(plaintext, decrypted) {
		t.Errorf("Round trip failed. Original: %v, Result: %v", plaintext, decrypted)
	}
}

