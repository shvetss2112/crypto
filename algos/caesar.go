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

func Caesar(isEncrypt bool, key string) func([]byte, []byte) {
	trueKey, err := strconv.Atoi(key)

	if err != nil {
		log.Fatal(err)
	}

	shift := byte(trueKey)

	if isEncrypt {
		return func(iBuf []byte, oBuf []byte) { caesarEnc(iBuf, oBuf, shift) }
	} else {
		return func(iBuf []byte, oBuf []byte) { caesarDec(iBuf, oBuf, shift) }
	}
}
