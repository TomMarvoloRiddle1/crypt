***HOW TO USE***

crypt/
├── internal/
│   └── app/
│       └── myapp.go # Core application logic
│   └── data/ #encryption storage
│       └──aes
│       └──plainText
│       └──rsa
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
