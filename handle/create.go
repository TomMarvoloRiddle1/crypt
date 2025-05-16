package handle

import (
	"fmt"
	"log"
	"os"
)

func CreateAndWriteFile(fileName string, text string) error {

	//sets working dir to "data" folder
	os.Chdir("data")

	textFileName := fmt.Sprintf("%s.txt", fileName)
	os.Create(textFileName)
	os.WriteFile(textFileName, []byte(text), 0666)

	return nil
}

func FileString(fileName string) string {
	rawTxtData, errReadFile := os.ReadFile(fileName)
	if errReadFile != nil {
		log.Fatal(errReadFile)
	}

	return string(rawTxtData)
}
