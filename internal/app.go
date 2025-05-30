package internal

import (
	"crypt/pkg"
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

	symmetricKey := pkg.AesKey()

	//rsa pub key encrypts the OG sym AES key
	rsaEncAesKey := pkg.RsaEnc(symmetricKey)
	os.WriteFile("./data/aes/rsaEncAesKey", rsaEncAesKey, 0666) //reverse this key for decryption, get symmetricKey

	pkg.AesEnc(fileName, symmetricKey)
}

func Decrypt(fileName string) {

	ogAesKey := pkg.DecAesWithRsa()

	pkg.DecText(fileName, ogAesKey)
}
