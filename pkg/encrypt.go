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

// GREEN
func Selection() (string, error) {
	//hardcoded dir, BAD when I scale
	folder, err := os.ReadDir("./data/plainText")
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

	fileSuffix := strings.TrimSuffix(txtFiles[indexList], ".txt")
	nameOnly := strings.TrimLeft(fileSuffix, "- ")

	return nameOnly, nil
}

func EntireEnc(plainTextName string) {

	originalDataName := fmt.Sprintf("./data/plainText/%s.txt", plainTextName)
	byteDataOg, _ := os.ReadFile(originalDataName)
	strDataOg := string(byteDataOg)

	key := make([]byte, 32)
	if _, err := rand.Reader.Read(key); err != nil {
		fmt.Println("error generating random encryption key ", err)
		return
	}

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

	ciphertext := gcm.Seal(nonce, nonce, []byte(strDataOg), nil)
	enc := hex.EncodeToString(ciphertext)

	// writing to pks and enc

	pkDir := fmt.Sprintf("./data/pks/%s", plainTextName)
	os.Create(pkDir)
	os.WriteFile(pkDir, key, 0666)

	encDir := fmt.Sprintf("./data/enc/%s_enc.txt", plainTextName)
	os.Create(encDir)
	os.WriteFile(encDir, []byte(enc), 0666)

}
