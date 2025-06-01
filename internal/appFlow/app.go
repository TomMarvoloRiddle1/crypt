package appFlow

import (
	"crypt/internal/decryption"
	"crypt/internal/encryption"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"log"
	"os"
)

func GenRsa() {
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

func Encrypt(fileName string) {

	symmetricKey := encryption.AesKey()

	//rsa pub key encrypts the OG sym AES key
	rsaEncAesKey := encryption.RsaEnc(symmetricKey)
	os.WriteFile("./data/aes/rsaEncAesKey", rsaEncAesKey, 0666) //reverse this key for decryption, get symmetricKey

	encryption.AesEnc(fileName, symmetricKey)
}

func Decrypt(fileName string) {

	ogAesKey := decryption.DecAesWithRsa()

	decryption.DecText(fileName, ogAesKey)
}
