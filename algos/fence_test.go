package algos

import (
	"bytes"
	"testing"
)

func TestFenceRoundTrip(t *testing.T) {
	plaintext := []byte("thequickbrownfox")
	key := "2"

	encFunc := Fence(true, key)
	decFunc := Fence(false, key)

	encrypted := make([]byte, len(plaintext))
	decrypted := make([]byte, len(plaintext))

	encFunc(plaintext, encrypted)
	decFunc(encrypted, decrypted)

	if !bytes.Equal(plaintext, decrypted) {
		t.Errorf("Round trip failed. Original: %v, Result: %v", plaintext, decrypted)
	}
}
