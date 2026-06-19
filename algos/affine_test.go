package algos

import (
	"bytes"
	"testing"
)

func TestAffineRoundTrip(t *testing.T) {
	plaintext := []byte("The quick brown fox")
	key := "3,5"

	encFunc := Affine(true, key)
	decFunc := Affine(false, key)

	encrypted := make([]byte, len(plaintext))
	decrypted := make([]byte, len(plaintext))

	encFunc(plaintext, encrypted)
	decFunc(encrypted, decrypted)

	if !bytes.Equal(plaintext, decrypted) {
		t.Errorf("Round trip failed. Original: %v, Result: %v", plaintext, decrypted)
	}
}

func TestModInverse(t *testing.T) {
	_, err := modInverse(3)
	if err != nil {
		t.Errorf("modInverse(3) failed: %v", err)
	}
}
