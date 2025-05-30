package main

import (
	"crypt/pkg"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"log"
	"os"
)

func main() {
	initModules(initSelect())
}

func initSelect() int {
	selectionOpt := `
1) Encrypt text file
2) Decrypt text file
3) Create RSA keypair
4) Exit program`
	fmt.Println(selectionOpt)
	var selectionUser int
	fmt.Scan(&selectionUser)
	return selectionUser
}

func initModules(selectionUser int) {

	switch selectionUser {
	case 1:
		readPlainTxt, errInStr := pkg.TargetNameEnc()
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

	case 2:

		target, _ := pkg.TargetNameDec()

		ogAesKey := pkg.DecAesWithRsa()

		pkg.DecText(target, ogAesKey)

	case 3:
		privK, privKeyErr := rsa.GenerateKey(rand.Reader, 4096)
		if privKeyErr != nil {
			log.Fatal(privKeyErr)
		}
		pubK := privK.PublicKey

		pubKeyByte := x509.MarshalPKCS1PublicKey(&pubK)
		privKeyByte := x509.MarshalPKCS1PrivateKey(privK)
		os.WriteFile("./data/rsa/priv", privKeyByte, 0666)
		os.WriteFile("./data/rsa/pub", pubKeyByte, 0666)

	case 4:
		os.Exit(3)

	default:
		fmt.Println("invalid selection")
		initModules(initSelect())
	}

}
