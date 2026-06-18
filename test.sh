go  build -o bin  cmd/crypto/main.go 
./bin encrypt fence data/ducati.jpg data/cipher 500 50
./bin decrypt fence data/cipher data/check.jpg 500 50