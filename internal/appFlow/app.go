package appFlow

import (
	"crypt/internal/decryption"
	"crypt/internal/encryption"
	"crypt/internal/zipping"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"log"
	"os"
	"regexp"
)

func GenRsa() {
	privK, privKeyErr := rsa.GenerateKey(rand.Reader, 4096)
	if privKeyErr != nil {
		log.Fatal(privKeyErr)
	}
	pubK := privK.PublicKey

	pubKeyByte := x509.MarshalPKCS1PublicKey(&pubK)
	privKeyByte := x509.MarshalPKCS1PrivateKey(privK)
	os.WriteFile("./data/keys/rsa/priv", privKeyByte, 0666)
	os.WriteFile("./data/keys/rsa/pub", pubKeyByte, 0666)
}

func Encrypt(fileName string) {

	symmetricKey := encryption.AesKey()

	//rsa pub key encrypts the OG sym AES key
	rsaEncAesKey := encryption.RsaEnc(symmetricKey)
	os.WriteFile("./data/keys/aes_enc/rsaEncAesKey", rsaEncAesKey, 0666) //reverse this key for decryption, get symmetricKey

	isZip, _ := regexp.MatchString(".zip", fileName)
	if isZip {
		srcPathZip := fmt.Sprintf("./data/exports/foldersZipped/%s", fileName)
		encryption.AesEnc(fileName, symmetricKey, srcPathZip)
	} else {
		srcPath := fmt.Sprintf("./data/imports/files/%s", fileName)
		encryption.AesEnc(fileName, symmetricKey, srcPath)
	}

}

func Decrypt(fileName string) {

	ogAesKey := decryption.DecAesWithRsa()

	srcPath := fmt.Sprintf("./data/exports/encrypted/%s_aesEnc", fileName)
	decryption.DecText(fileName, ogAesKey, srcPath)

}

func Zip(folderName string) {

	originFolder := fmt.Sprintf("./data/imports/folders/%s", folderName)
	exportedZip := fmt.Sprintf("./data/exports/foldersZipped/%s.zip", folderName)

	zipping.ZipTargetFolder(originFolder, exportedZip)

}
