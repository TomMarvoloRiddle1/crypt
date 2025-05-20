package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func test() {
	fmt.Println("test")
}

func DecText() {

	key, _ := os.ReadFile("./data/pks/one")

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("error creating aes block cipher", err)
		return
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("error setting gcm mode", err)
		return
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println("error generating the nonce ", err)
		return
	}

	a, _ := os.ReadFile("./data/enc/one_enc.txt")
	enc := string(a)

	decodedCipherText, err := hex.DecodeString(enc)
	if err != nil {
		fmt.Println("error decoding hex", err)
		return
	}

	decryptedData, err := gcm.Open(nil, decodedCipherText[:gcm.NonceSize()], decodedCipherText[gcm.NonceSize():], nil)
	if err != nil {
		fmt.Println("error decrypting data", err)
		return
	}

	fmt.Println("Decrypted data:", string(decryptedData))

}
