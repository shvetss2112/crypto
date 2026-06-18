go  build -o bin  cmd/crypto/main.go

test(){
echo Cipher: $1 
echo Key: $2

echo Encrypt with chunk size 10
./bin encrypt $1 data/ducati.jpg data/cipher 10 $2
echo 
echo Encrypt with chunk size 500
./bin encrypt $1 data/ducati.jpg data/cipher 500 $2
echo
echo Encrypt with chunk size 10000
./bin encrypt $1 data/ducati.jpg data/cipher 10000 $2
echo 
echo Encrypt with chunk size 6000000
./bin encrypt $1 data/ducati.jpg data/cipher_sequential 6000000 $2

echo
echo Decrypt
./bin decrypt $1 data/cipher data/check.jpg 10000 $2
}

test caesar 200

