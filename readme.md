***HOW TO USE***

crypt/
├── data/ #encryption storage
│    └── app/
│       └── main.go #entry point, where program is build from
├── data/ #user resources, must make directories yourself
│    └──exports
│    └──imports
│    └──aes
│    └──plainText
│    └──rsa
├── internal/
│   └── myapp.go #module logic
├── pkg/ #reusable functions
│   └── somepkg/
│       └── somepkg.go # Reusable functions
├── .gitignore
├── go.mod
├── main.go
├── README.md

plainText\ contains plaintext (.txt)
enc\ contains encrypted files +
dec\ destinated of decrypted text
pks\ where private keys stored to be read from


build native:
go build -trimpath -o crypt main.go
build win:
GOOS=windows GOARCH=amd64 go build -trimpath -o crypt.exe main.go
build mac:
GOOS=darwin GOARCH=arm64 go build -trimpath -o crypt main.go


usage after building:
./crypt genrsa
- creates rsa keypair under ./data/rsa
./crypt encrypt <NAME>
- encrypts message inside of .txt file, DO NOT INCLUDE .txt JUST THE TEXT FILE NAME
./crypt decrypt <NAME>
- decrypts message inside of the aes encrypted file from hexadecimal string #need to refine this


notes:
<NAME> is volatile rn, since it changes to a string format for <name>_aesEnc.txt, reads native <name>.txt
lots of hardcoded logic like rsaEncAesKey in aes (the encrypted AES key via rsaPub, need RSA PK to access)
make dynamic naming of keys to handle more files, rn just a MVP to streamline logic
- zero streamlining inside encrypt decrypt modules, also alot of omitted error handling bcoz lazy (but some of it is truly unnecessary)

>> could storing the keys in a .json be more efficient?
