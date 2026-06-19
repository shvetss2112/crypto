package algos

import (
	"bytes"
	"testing"
)

func TestAffineEncrypt(t *testing.T) {
	tests := []struct {
		name      string
		plaintext []byte
		key       string
	}{
		{
			name:      "key 3,5",
			plaintext: []byte("hello"),
			key:       "3,5",
		},
		{
			name:      "key 1,0 identity",
			plaintext: []byte("test"),
			key:       "1,0",
		},
		{
			name:      "single byte",
			plaintext: []byte("a"),
			key:       "5,3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encFunc := Affine(true, tt.key)
			output := make([]byte, len(tt.plaintext))
			encFunc(tt.plaintext, output)

			if len(output) != len(tt.plaintext) {
				t.Errorf("Output length mismatch. Got %d, expected %d", len(output), len(tt.plaintext))
			}
		})
	}
}

func TestAffineDecrypt(t *testing.T) {
	tests := []struct {
		name       string
		ciphertext []byte
		key        string
	}{
		{
			name:       "key 1,0 identity roundtrip",
			ciphertext: []byte("test"),
			key:        "1,0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decFunc := Affine(false, tt.key)
			output := make([]byte, len(tt.ciphertext))
			decFunc(tt.ciphertext, output)

			if len(output) != len(tt.ciphertext) {
				t.Errorf("Output length mismatch. Got %d, expected %d", len(output), len(tt.ciphertext))
			}
		})
	}
}

func TestAffineRoundTrip(t *testing.T) {
	plaintext := []byte("The quick brown fox")
	keys := []string{"3,5", "5,7", "7,11", "9,13", "11,17"}

	for _, key := range keys {
		t.Run("roundtrip_key_"+key, func(t *testing.T) {
			encFunc := Affine(true, key)
			decFunc := Affine(false, key)

			encrypted := make([]byte, len(plaintext))
			decrypted := make([]byte, len(plaintext))

			encFunc(plaintext, encrypted)
			decFunc(encrypted, decrypted)

			if !bytes.Equal(plaintext, decrypted) {
				t.Errorf("Round trip failed with key %s. Original: %v, Result: %v", key, plaintext, decrypted)
			}
		})
	}
}


func TestAffineByteBoundaries(t *testing.T) {
	plaintext := []byte{0, 127, 255, 1, 128}
	key := "3,5"

	encFunc := Affine(true, key)
	decFunc := Affine(false, key)

	encrypted := make([]byte, len(plaintext))
	decrypted := make([]byte, len(plaintext))

	encFunc(plaintext, encrypted)
	decFunc(encrypted, decrypted)

	if !bytes.Equal(plaintext, decrypted) {
		t.Errorf("Byte boundary test failed. Original: %v, Result: %v", plaintext, decrypted)
	}
}

func TestModInverse(t *testing.T) {
	tests := []struct {
		name  string
		a     byte
		valid bool
	}{
		{
			name:  "inverse of 3",
			a:     3,
			valid: true,
		},
		{
			name:  "inverse of 5",
			a:     5,
			valid: true,
		},
		{
			name:  "inverse of 7",
			a:     7,
			valid: true,
		},
		{
			name:  "even number has no inverse",
			a:     2,
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inv, err := modInverse(tt.a)
			if tt.valid {
				if err != nil {
					t.Errorf("Expected valid inverse for %d, got error: %v", tt.a, err)
				}
				// Verify the inverse is correct: (a * inv) mod 256 == 1
				result := (int(tt.a) * int(inv)) % 256
				if result != 1 {
					t.Errorf("Invalid inverse. %d * %d mod 256 = %d, expected 1", tt.a, inv, result)
				}
			} else {
				if err == nil {
					t.Errorf("Expected error for %d, got inverse: %d", tt.a, inv)
				}
			}
		})
	}
}
