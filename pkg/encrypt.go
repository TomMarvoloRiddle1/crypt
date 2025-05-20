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
	return txtFiles[indexList], nil

}

func EncKeyOne(fileName string) ([]byte, string) {
	aesKey := make([]byte, 32)

	if _, errKey := rand.Reader.Read(aesKey); errKey != nil {
		log.Fatal(errKey)
	}

	fileSuffix := strings.TrimSuffix(fileName, ".txt")
	fileNom := strings.TrimLeft(fileSuffix, "- ")

	//fileName passes numbers and old string formatting
	encKeyFile := fmt.Sprintf("./data/pks/%s_key", fileNom)

	os.Create(encKeyFile)
	//byte data written in second param
	os.WriteFile(encKeyFile, aesKey, 0666)

	// under ~/data/FILENAME_key a 32bit key is generated
	return aesKey, encKeyFile

}

func EncProcTwo(fileName string, pk []byte) []byte {

	textData, _ := os.ReadFile(fileName)

	//pk aka privateKey used to determine Data block here
	block, errBlock := aes.NewCipher(pk)
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

	cipherData := gcm.Seal(nonce, nonce, textData, nil)

	hexString := hex.EncodeToString(cipherData)

	return []byte(hexString)

}

func EncWrite(cipherData []byte, baseName string) {

	baseNameTwo := strings.ReplaceAll(baseName, "./data/pks/", "")
	encKeyFile := fmt.Sprintf("./data/enc/%s_enc.txt", baseNameTwo)

	os.Create(encKeyFile)
	//byte data written in second param
	os.WriteFile(encKeyFile, cipherData, 0666)
}
