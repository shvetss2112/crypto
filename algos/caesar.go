package algos

import (
	"log"
	"strconv"
)

func caesarEnc(iBuf []byte, oBuf []byte, shift byte) {
	size := len(iBuf)

	for i := 0; i < size; i++ {
		oBuf[i] = iBuf[i] + shift
	}
}

func caesarDec(iBuf []byte, oBuf []byte, shift byte) {
	size := len(iBuf)

	for i := 0; i < size; i++ {
		oBuf[i] = iBuf[i] - shift
	}
}

func Caesar(iBuf []byte, oBuf []byte, isEncrypt bool, key string) {
	trueKey, err := strconv.Atoi(key)

	if err != nil {
		log.Fatal(err)
	}

	shift := byte(trueKey)

	if isEncrypt {
		caesarEnc(iBuf, oBuf, shift)
	} else {
		caesarDec(iBuf, oBuf, shift)
	}
}
