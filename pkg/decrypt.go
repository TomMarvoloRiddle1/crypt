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
	"regexp"
	"strings"
)

func SelectionDec() (string, error) {
	//hardcoded dir, BAD when I scale
	folder, err := os.ReadDir("./data/enc")
	if err != nil {
		fmt.Println("dir issue 1")
		log.Fatal(err)
	}

	var txtFiles []string
	for i, v := range folder {
		currNum := i + 1
		availFiles := fmt.Sprintf("%d) %v", currNum, v)
		currFileName := fmt.Sprintf("%v", v)

		// check only for text files
		isTxt, _ := regexp.MatchString(".txt", availFiles)
		if isTxt {
			fmt.Println(availFiles)
			txtFiles = append(txtFiles, currFileName)
		} else {
			continue
		}

	}

	var selectedFile int
	fmt.Println("which file would you like to encrypt?")
	fmt.Scan(&selectedFile)

	indexList := selectedFile - 1

	fileSuffix := strings.TrimSuffix(txtFiles[indexList], "_enc.txt")
	nameOnly := strings.ReplaceAll(fileSuffix, "- ", "")

	return nameOnly, nil
}

func DecText(target string) {

	keyTarg := fmt.Sprintf("./data/pks/%s", target)
	key, _ := os.ReadFile(keyTarg)

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

	encTarg := fmt.Sprintf("./data/enc/%s_enc.txt", target)
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

	decFile := fmt.Sprintf("./data/dec/%s_dec.txt", target)
	os.Create(decFile)
	os.WriteFile(decFile, decryptedData, 0666)

	fmt.Println("Decrypted data:", string(decryptedData))

}
