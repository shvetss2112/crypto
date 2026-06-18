package main

import (
	algos "crypto/algos"
	"log"
	"os"
	"sync"
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
	case "copy":
		return func(iBuf []byte, oBuf []byte) { algos.Copy(iBuf, oBuf, config.isEncrypt, config.strKey) }
	case "fence":
		return func(iBuf []byte, oBuf []byte) { algos.Fence(iBuf, oBuf, config.isEncrypt, config.strKey) }
	case "kardano":
		return func(iBuf []byte, oBuf []byte) { algos.Copy(iBuf, oBuf, config.isEncrypt, config.strKey) }
	case "some":
		return func(iBuf []byte, oBuf []byte) { algos.Copy(iBuf, oBuf, config.isEncrypt, config.strKey) }
	case "other":
		return func(iBuf []byte, oBuf []byte) { algos.Copy(iBuf, oBuf, config.isEncrypt, config.strKey) }
	case "shit":
		return func(iBuf []byte, oBuf []byte) { algos.Copy(iBuf, oBuf, config.isEncrypt, config.strKey) }
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
	iFile, oFile := manageFiles(config.iFileName, config.oFileName)
	defer oFile.Close()
	defer iFile.Close()
	ParalelStart(iFile, oFile, config.chunkSize, algoChooser(config))
}

func main() {

	config := &cryptoConfig{
		iFileName: "data/ducati.jpg",
		oFileName: "data/cipher",
		algo:      "fence",
		chunkSize: 400,
		strKey:    "20",
		isEncrypt: true,
	}

	run(config)
	config2 := &cryptoConfig{
		iFileName: "data/cipher",
		oFileName: "data/check.jpg",
		algo:      "fence",
		chunkSize: 400,
		strKey:    "20",
		isEncrypt: false,
	}
	run(config2)
}
