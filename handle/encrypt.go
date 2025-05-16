package handle

import (
	"crypto/rand"
	"fmt"
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

func EncKey(fileName string) []byte {
	aesKey := make([]byte, 32)

	if _, errKey := rand.Reader.Read(aesKey); errKey != nil {
		log.Fatal(errKey)
	}

	fileSuffix := strings.TrimSuffix(fileName, ".txt")
	fileNom := strings.TrimLeft(fileSuffix, "- ")

	//fileName passes numbers and old string formatting
	encKeyFile := fmt.Sprintf("%s_key.txt", fileNom)
	os.Chdir("data")
	os.Create(encKeyFile)
	os.WriteFile(encKeyFile, aesKey, 0666)

	return aesKey

}
