package algos

// responcibility: give callback func([]byte, []byte) for particular algo

import (
	"log"
	"strconv"
)

func fenceEnc(iBuf []byte, oBuf []byte, key int) {
	size := len(iBuf)

	counter := 0
	for offset := 0; offset < key; offset++ {
		for i := key - 1; i < size+key; i += 2*key - 2 {
			if i-offset < size && i-offset >= 0 {
				oBuf[counter] = iBuf[i-offset]
				counter++
			}
			if offset != 0 && i+offset < size && offset < key-1 {
				oBuf[counter] = iBuf[i+offset]
				counter++
			}
		}
	}
}

func fenceDec(iBuf []byte, oBuf []byte, key int) {
	size := len(iBuf)

	counter := 0
	for offset := 0; offset < key; offset++ {
		for i := key - 1; i < size+key; i += 2*key - 2 {
			if i-offset < size && i-offset >= 0 {
				oBuf[i-offset] = iBuf[counter]
				counter++
			}
			if offset != 0 && i+offset < size && offset < key-1 {
				oBuf[i+offset] = iBuf[counter]
				counter++
			}
		}
	}
}

func Fence(isEncrypt bool, key string) func([]byte, []byte) {
	trueKey, err := strconv.Atoi(key)

	if err != nil {
		log.Fatal(err)
	}

	if isEncrypt {
		return func(iBuf []byte, oBuf []byte) { fenceEnc(iBuf, oBuf, trueKey) }
	} else {
		return func(iBuf []byte, oBuf []byte) { fenceDec(iBuf, oBuf, trueKey) }
	}

}
