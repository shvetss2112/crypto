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

func Vigenere(iBuf []byte, oBuf []byte, isEncrypt bool, key string) {
	if len(key) == 0 {
		log.Fatal("vigenere key cannot be empty")
	}

	keyBytes := []byte(key)

	if isEncrypt {
		vigenereEnc(iBuf, oBuf, keyBytes)
	} else {
		vigenereDec(iBuf, oBuf, keyBytes)
	}
}
