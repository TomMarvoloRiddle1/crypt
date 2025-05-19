package handle

import (
	"fmt"
	"regexp"
)

func ReadDataFolderDec() []string {
	//hardcoded dir, BAD when I scale
	dataList := checkDir("data")

	var txtFiles []string
	for i, v := range dataList {
		currNum := i + 1
		availFiles := fmt.Sprintf("%d) %v", currNum, v)
		currFileName := fmt.Sprintf("%v", v)

		// check only for text files
		isTxt, _ := regexp.MatchString("enc", availFiles)
		if isTxt != true {
			continue
		} else {
			fmt.Println(availFiles)
			txtFiles = append(txtFiles, currFileName)
		}

	}
	return txtFiles

}

func DecryptMod(hexEnc []byte) {

}
