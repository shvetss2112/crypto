package algos

import (
	"bytes"
	"testing"
)

func TestCaesarRoundTrip(t *testing.T) {
	plaintext := []byte("The quick brown fox")
	shift := "3"

	encFunc := Caesar(true, shift)
	decFunc := Caesar(false, shift)

	encrypted := make([]byte, len(plaintext))
	decrypted := make([]byte, len(plaintext))

	encFunc(plaintext, encrypted)
	decFunc(encrypted, decrypted)

	if !bytes.Equal(plaintext, decrypted) {
		t.Errorf("Round trip failed. Original: %v, Result: %v", plaintext, decrypted)
	}
}

