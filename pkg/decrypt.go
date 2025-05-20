package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func RetrieveGCM() cipher.AEAD {
	data, _ := os.ReadFile("./data/pks/three_key")

	block, errBlock := aes.NewCipher(data)
	if errBlock != nil {
		log.Fatal(errBlock)
	}

	gcm, errGcm := cipher.NewGCM(block)
	if errGcm != nil {
		log.Fatal(errGcm)
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, errNo := io.ReadFull(rand.Reader, nonce); errNo != nil {
		log.Fatal(errNo)
	}

	return gcm
}

func DecFlow() {

	gcm := RetrieveGCM()

	encByte, _ := os.ReadFile("./data/enc/three_key_enc.txt")
	data := string(encByte)

	// fmt.Println(data)

	decodedCipherText, _ := hex.DecodeString(data)

	decStr, _ := gcm.Open(nil, decodedCipherText[:gcm.NonceSize()], decodedCipherText[gcm.NonceSize():], nil)

	fmt.Println(string(decStr))

}
