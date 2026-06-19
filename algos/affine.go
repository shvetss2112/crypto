package algos

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func egcd(a, b int) (int, int, int) {
	if a == 0 {
		return b, 0, 1
	}

	gcd, x1, y1 := egcd(b%a, a)
	x := y1 - (b/a)*x1
	y := x1

	return gcd, x, y
}

func modInverse(a byte) (byte, error) {
	gcd, x, _ := egcd(int(a), 256)

	if gcd != 1 {
		return 0, fmt.Errorf("no inverse")
	}

	res := x % 256
	if res < 0 {
		res += 256
	}

	return byte(res), nil
}

func affineEnc(iBuf []byte, oBuf []byte, a, b byte) {
	size := len(iBuf)
	for i := 0; i < size; i++ {
		oBuf[i] = a*iBuf[i] + b
	}
}

func affineDec(iBuf []byte, oBuf []byte, aInv, b byte) {
	size := len(iBuf)
	for i := 0; i < size; i++ {
		oBuf[i] = aInv * (iBuf[i] - b)
	}
}

func Affine(isEncrypt bool, key string) func([]byte, []byte) {
	parts := strings.Split(key, ",")
	if len(parts) != 2 {
		log.Fatal("Invalid key format")
	}

	aInt, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		log.Fatal("Invalid key format")
	}

	bInt, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		log.Fatal("Invalid key format")
	}

	a := byte(aInt)
	b := byte(bInt)

	if a%2 == 0 {
		log.Fatal("a must be odd")
	}

	if isEncrypt {
		return func(iBuf []byte, oBuf []byte) { affineEnc(iBuf, oBuf, a, b) }
	}

	aInv, err := modInverse(a)
	if err != nil {
		log.Fatal(err)
	}

	return func(iBuf []byte, oBuf []byte) { affineDec(iBuf, oBuf, aInv, b) }
}
