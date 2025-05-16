package handle

import (
	"fmt"
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
