package pkg

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func TargetNameDec() (string, error) {

	fmt.Println("enter name of file to decrypt [without .txt]")
	var namePlainTxt string
	fmt.Scan(&namePlainTxt)

	return namePlainTxt, nil
}

func DecAesWithRsa() []byte {
	//consider filepath as parameter

	aesKeyEncWithRsa, _ := os.ReadFile("./aes/rsaEncAesKey")

	priv, _ := os.ReadFile("./rsa/priv")

	privStructure, _ := x509.ParsePKCS1PrivateKey(priv)
	ogAesKey, _ := privStructure.Decrypt(nil, aesKeyEncWithRsa, &rsa.OAEPOptions{Hash: crypto.SHA256})

	return ogAesKey
}

func DecText(target string, ogAesKey []byte) {

	block, err := aes.NewCipher(ogAesKey)
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

	//hybrid encrypted shit
	encTarg := fmt.Sprintf("./aes/%s_aesEnc.txt", target)
	encByte, _ := os.ReadFile(encTarg)
	enc := string(encByte)

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
