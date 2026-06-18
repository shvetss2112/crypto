package algos

import (
	"log"
)

func autokeyEnc(iBuf []byte, oBuf []byte, keyBytes []byte) {
	keyLen := len(keyBytes)
	size := len(iBuf)

	for i := 0; i < size; i++ {
		var k byte
		if i < keyLen {
			k = keyBytes[i]
		} else {
			k = iBuf[i-keyLen]
		}
		oBuf[i] = iBuf[i] + k
	}
}

func autokeyDec(iBuf []byte, oBuf []byte, keyBytes []byte) {
	keyLen := len(keyBytes)
	size := len(iBuf)

	for i := 0; i < size; i++ {
		var k byte
		if i < keyLen {
			k = keyBytes[i]
		} else {
			k = oBuf[i-keyLen]
		}
		oBuf[i] = iBuf[i] - k
	}
}

func Autokey(isEncrypt bool, key string) func([]byte, []byte) {
	if len(key) == 0 {
		log.Fatal("autokey key cannot be empty")
	}

	keyBytes := []byte(key)

	if isEncrypt {
		return func(iBuf []byte, oBuf []byte) { autokeyEnc(iBuf, oBuf, keyBytes) }
	} else {
		return func(iBuf []byte, oBuf []byte) { autokeyDec(iBuf, oBuf, keyBytes) }
	}
}
