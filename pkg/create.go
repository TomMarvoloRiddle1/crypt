package pkg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// step 1, takes filename + text data as inputs
func PromptData() (string, string) {
	oneSelectName := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text file name: ")
	//file name
	fileNameInRaw, errFileName := oneSelectName.ReadString('\n')
	if errFileName != nil {
		log.Fatal(errFileName)
	}
	fileNameIn := strings.TrimSpace(fileNameInRaw)

	oneSelectBody := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text file info/data: ")
	//text data
	fileTextIn, errTextIn := oneSelectBody.ReadString('\n')
	if errTextIn != nil {
		fmt.Println("invalid entry, likely not a string or file corruption")
		log.Fatal(errTextIn)
	}

	return fileNameIn, fileTextIn
}

// step2
func CreateAndWriteFile(fileName string, text string) error {

	textFileName := fmt.Sprintf("./data/plainText/%s.txt", fileName)
	os.Create(textFileName)
	os.WriteFile(textFileName, []byte(text), 0666)

	return nil
}
