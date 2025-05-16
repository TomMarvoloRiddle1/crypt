package handle

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func ReadDataFolder() {
	dataList := checkDir("data")
	for i, v := range dataList {
		currNum := i + 1
		availFiles := fmt.Sprintf("%d) %v\n", currNum, v)

		isTxt, _ := regexp.MatchString(".txt", availFiles)
		if isTxt != true {
			continue
		} else {
			fmt.Println(availFiles)
		}

	}

}

func checkDir(folderName string) []os.DirEntry {
	folder, err := os.ReadDir(folderName)
	if err != nil {
		log.Fatal(err)
	}

	return folder
}
