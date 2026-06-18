package algos

func copyEnc(iBuf []byte, oBuf []byte, key string) {
	copy(oBuf, iBuf)
}

func copyDec(iBuf []byte, oBuf []byte, key string) {
	copy(oBuf, iBuf)
}

func Copy(iBuf []byte, oBuf []byte, isEncrypt bool, key string) {
	if isEncrypt {
		copyEnc(iBuf, oBuf, key)
	} else {
		copyEnc(iBuf, oBuf, key)
	}
}
