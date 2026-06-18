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
	chunkSize int
	strKey    string
	isEncrypt bool
}

func manageFiles(iFileName string, oFileName string) (*os.File, *os.File) {
	//open input file this one
	iFile, err := os.Open(iFileName)
	if err != nil {
		log.Fatal(err)
	}

	// create output file
	oFile, err := os.Create(oFileName)
	if err != nil {
		log.Fatal(err)
	}

	return iFile, oFile

}

func algoChooser(config *cryptoConfig) func(iBuf []byte, oBuf []byte) {

	switch config.algo {
	case "fence":
		return func(iBuf []byte, oBuf []byte) { algos.Fence(iBuf, oBuf, config.isEncrypt, config.strKey) }
	case "kardano":
		return func(iBuf []byte, oBuf []byte) { algos.Fence(iBuf, oBuf, config.isEncrypt, config.strKey) }
	case "vizhener":
		return func(iBuf []byte, oBuf []byte) { algos.Fence(iBuf, oBuf, config.isEncrypt, config.strKey) }
	case "something":
		return func(iBuf []byte, oBuf []byte) { algos.Fence(iBuf, oBuf, config.isEncrypt, config.strKey) }
	case "else":
		return func(iBuf []byte, oBuf []byte) { algos.Fence(iBuf, oBuf, config.isEncrypt, config.strKey) }
	default:
		log.Fatal("Unexpected algorithm: " + config.algo)
		return nil
	}
}

func ParalelStart(iFile *os.File, oFile *os.File, blockSize int, algo func([]byte, []byte)) {

	// get filesize
	fstat, err := iFile.Stat()
	if err != nil {
		log.Fatal(err)
	}
	var fsz int64 = fstat.Size()

	//spin ts up
	var wg sync.WaitGroup
	for i := int64(0); i < fsz; i = i + int64(blockSize) {
		iBuffer := make([]byte, blockSize)
		oBuffer := make([]byte, blockSize)
		iFile.ReadAt(iBuffer, i)
		wg.Go(func() { algo(iBuffer, oBuffer); oFile.WriteAt(oBuffer, i) })
	}
	wg.Wait()
}

func run(config *cryptoConfig) {
	start := time.Now()
	iFile, oFile := manageFiles(config.iFileName, config.oFileName)
	defer oFile.Close()
	defer iFile.Close()
	ParalelStart(iFile, oFile, config.chunkSize, algoChooser(config))
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}

func main() {

	var actualIsEncrypt bool

	if os.Args[1] == "encrypt" {
		actualIsEncrypt = true
	} else if os.Args[1] == "decrypt" {
		actualIsEncrypt = false
	} else {
		log.Fatal("Unknown command: " + os.Args[1])
	}
	actualChunkSize, err := strconv.Atoi(os.Args[5])
	if err != nil {
		log.Fatal("Incorrenct argument: " + os.Args[5])
	}
	config := &cryptoConfig{
		isEncrypt: actualIsEncrypt,
		algo:      os.Args[2],
		iFileName: os.Args[3],
		oFileName: os.Args[4],
		chunkSize: actualChunkSize,
		strKey:    os.Args[6],
	}

	run(config)
}
