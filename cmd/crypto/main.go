package main

import (
	algos "crypto/algos"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

type cryptoConfig struct {
	iFileName string
	oFileName string
	algo      string
	threads   int
	strKey    string
	isEncrypt bool
}

func getFiles(iFileName string, oFileName string) (*os.File, *os.File) {
	iFile, err := os.Open(iFileName)
	if err != nil {
		log.Fatal(err)
	}

	oFile, err := os.Create(oFileName)
	if err != nil {
		log.Fatal(err)
	}

	return iFile, oFile

}

func chooseAlgo(config *cryptoConfig) func(iBuf []byte, oBuf []byte) {
	// resp: decide which particular callback from all
	switch config.algo {
	case "fence":
		return algos.Fence(config.isEncrypt, config.strKey)
	case "caesar":
		return algos.Caesar(config.isEncrypt, config.strKey)
	case "vigenere":
		return algos.Vigenere(config.isEncrypt, config.strKey)
	case "something":
		return algos.Fence(config.isEncrypt, config.strKey)
	default:
		log.Fatal("Unexpected algorithm: " + config.algo)
		return nil
	}
}

func startInParalel(iFile *os.File, oFile *os.File, threadsNum int, algo func([]byte, []byte)) {

	fstat, err := iFile.Stat()
	if err != nil {
		log.Fatal(err)
	}
	var fsz int64 = fstat.Size()
	blockSize := (fsz + int64(threadsNum) - 1) / int64(threadsNum)
	if int64(threadsNum) > fsz || threadsNum < 1 {
		log.Fatal("Invalid number of threads")
	}

	var wg sync.WaitGroup
	for i := int64(0); i < fsz; i = i + int64(blockSize) {
		iBuffer := make([]byte, blockSize)
		oBuffer := make([]byte, blockSize)
		iFile.ReadAt(iBuffer, i)
		wg.Go(func() { algo(iBuffer, oBuffer); oFile.WriteAt(oBuffer, i) })
	}
	wg.Wait()
}

func config() *cryptoConfig {
	var actualIsEncrypt bool

	if os.Args[1] == "encrypt" {
		actualIsEncrypt = true
	} else if os.Args[1] == "decrypt" {
		actualIsEncrypt = false
	} else {
		log.Fatal("Unknown command: " + os.Args[1])
	}
	actualThreadCount, err := strconv.Atoi(os.Args[5])
	if err != nil {
		log.Fatal("Incorrenct argument: " + os.Args[5])
	}
	return &cryptoConfig{
		isEncrypt: actualIsEncrypt,
		algo:      os.Args[2],
		iFileName: os.Args[3],
		oFileName: os.Args[4],
		threads:   actualThreadCount,
		strKey:    os.Args[6],
	}
}

func run(config *cryptoConfig) {
	start := time.Now()
	iFile, oFile := getFiles(config.iFileName, config.oFileName)
	defer oFile.Close()
	defer iFile.Close()
	startInParalel(iFile, oFile, config.threads, chooseAlgo(config))
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}

func main() {
	run(config())
}
