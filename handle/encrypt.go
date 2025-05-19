package handle

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

func ReadDataFolder() []string {
	//hardcoded dir, BAD when I scale
	dataList := checkDir("data")

	var txtFiles []string
	for i, v := range dataList {
		currNum := i + 1
		availFiles := fmt.Sprintf("%d) %v", currNum, v)
		currFileName := fmt.Sprintf("%v", v)

		// check only for text files
		isTxt, _ := regexp.MatchString(".txt", availFiles)
		if isTxt != true {
			continue
		} else {
			fmt.Println(availFiles)
			txtFiles = append(txtFiles, currFileName)
		}

	}
	return txtFiles

}

func checkDir(folderName string) []os.DirEntry {
	folder, err := os.ReadDir(folderName)
	if err != nil {
		log.Fatal(err)
	}

	return folder
}

func Selection(fileList []string) (string, error) {

	var selectedFile int
	fmt.Println("which file would you like to encrypt?")
	fmt.Scan(&selectedFile)

	indexList := selectedFile - 1

	return fileList[indexList], nil

}

func EncKeyOne(fileName string) ([]byte, string) {
	aesKey := make([]byte, 32)

	if _, errKey := rand.Reader.Read(aesKey); errKey != nil {
		log.Fatal(errKey)
	}

	fileSuffix := strings.TrimSuffix(fileName, ".txt")
	fileNom := strings.TrimLeft(fileSuffix, "- ")

	//fileName passes numbers and old string formatting
	encKeyFile := fmt.Sprintf("%s_key", fileNom)
	os.Chdir("data")
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

	encKeyFile := fmt.Sprintf("%s_enc", baseName)
	os.Chdir("data")
	os.Create(encKeyFile)
	//byte data written in second param
	os.WriteFile(encKeyFile, cipherData, 0666)
}
