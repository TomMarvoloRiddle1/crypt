package internal

import (
	"crypt/pkg"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"log"
	"os"
)

func ModOne() {
	readPlainTxt, errInStr := pkg.NameInput("encrypt")
	if errInStr != nil {
		fmt.Println("not a string")
		log.Fatal(errInStr)
		//add error handling for invalid entries, because we are ditching a list
	}

	symmetricKey := pkg.AesKey()

	//rsa pub key encrypts the OG sym AES key
	rsaEncAesKey := pkg.RsaEnc(symmetricKey)
	os.WriteFile("./data/aes/rsaEncAesKey", rsaEncAesKey, 0666) //reverse this key for decryption, get symmetricKey

	pkg.AesEnc(readPlainTxt, symmetricKey)
}

func ModTwo() {
	target, _ := pkg.NameInput("decrypt")

	ogAesKey := pkg.DecAesWithRsa()

	pkg.DecText(target, ogAesKey)
}

func ModThree() {
	privK, privKeyErr := rsa.GenerateKey(rand.Reader, 4096)
	if privKeyErr != nil {
		log.Fatal(privKeyErr)
	}
	pubK := privK.PublicKey

	pubKeyByte := x509.MarshalPKCS1PublicKey(&pubK)
	privKeyByte := x509.MarshalPKCS1PrivateKey(privK)
	os.WriteFile("./data/rsa/priv", privKeyByte, 0666)
	os.WriteFile("./data/rsa/pub", pubKeyByte, 0666)
}
