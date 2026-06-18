package algos

import (
	"log"
)

func vigenereEnc(iBuf []byte, oBuf []byte, keyBytes []byte) {
	keyLen := len(keyBytes)
	size := len(iBuf)

	for i := 0; i < size; i++ {
		oBuf[i] = iBuf[i] + keyBytes[i%keyLen]
	}
}

func vigenereDec(iBuf []byte, oBuf []byte, keyBytes []byte) {
	keyLen := len(keyBytes)
	size := len(iBuf)

	for i := 0; i < size; i++ {
		oBuf[i] = iBuf[i] - keyBytes[i%keyLen]
	}
}

func Vigenere(isEncrypt bool, key string) func([]byte, []byte) {
	if len(key) == 0 {
		log.Fatal("vigenere key cannot be empty")
	}

	keyBytes := []byte(key)

	if isEncrypt {
		return func(iBuf []byte, oBuf []byte) { vigenereEnc(iBuf, oBuf, keyBytes) }
	} else {
		return func(iBuf []byte, oBuf []byte) { vigenereDec(iBuf, oBuf, keyBytes) }
	}
}
