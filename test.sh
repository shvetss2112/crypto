go  build -o bin  cmd/crypto/main.go

test(){
echo Cipher: $1 
echo Key: $2

echo Encrypt with $2 goroutine\(s\)
./bin encrypt $1 data/crypto.txt results/cipher $2 $3
echo
echo Decrypt
./bin decrypt $1 results/cipher results/check.txt  $2 $3
}

test $1 $2 $3
