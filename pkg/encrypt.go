package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func TargetNameEnc() (string, error) {

	fmt.Println("enter name of file to encrypt [without .txt]")
	var namePlainTxt string
	fmt.Scan(&namePlainTxt)

	return namePlainTxt, nil
}

func RsaEnc(data []byte) []byte {

	pubByte, _ := os.ReadFile("./data/rsa/pub")

	pubStructure, _ := x509.ParsePKCS1PublicKey(pubByte)

	rsaEncByte, _ := rsa.EncryptOAEP(sha256.New(), rand.Reader, pubStructure, data, nil)

	return rsaEncByte

}

func AesKey() []byte {
	key := make([]byte, 32)
	if _, err := rand.Reader.Read(key); err != nil {
		fmt.Println("error generating random encryption key ", err)
	}
	return key
}

func AesEnc(plainTextName string, aesKey []byte) {

	originalDataName := fmt.Sprintf("./data/plainText/%s.txt", plainTextName)
	byteDataOg, _ := os.ReadFile(originalDataName)
	strDataOg := string(byteDataOg)

	//crucial to be passed

	block, err := aes.NewCipher(aesKey)
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

	ciphertext := gcm.Seal(nonce, nonce, []byte(strDataOg), nil)
	enc := hex.EncodeToString(ciphertext)

	encDir := fmt.Sprintf("./data/aes/%s_aesEnc.txt", plainTextName)
	os.Create(encDir)
	os.WriteFile(encDir, []byte(enc), 0666)
}
