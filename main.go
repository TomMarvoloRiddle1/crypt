package main

import (
	"bufio"
	"crypt/handle"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	initModules()
}

func initModules() {

	selectionOpt := `1) Create and write to text file
2) Encrypt text file
3) Decrypt text file
4) Exit program`

	fmt.Println(selectionOpt)

	var selectionUser int

	fmt.Scan(&selectionUser)

	// add "return to menus" later
	switch selectionUser {
	case 1:

		oneSelectName := bufio.NewReader(os.Stdin)
		fmt.Println("Enter text file name: ")
		fileNameInRaw, errFileName := oneSelectName.ReadString('\n')
		if errFileName != nil {
			log.Fatal(errFileName)
		}
		fileNameIn := strings.TrimSpace(fileNameInRaw)

		oneSelectBody := bufio.NewReader(os.Stdin)
		fmt.Println("Enter text file info/data: ")
		fileTextIn, errTextIn := oneSelectBody.ReadString('\n')
		if errTextIn != nil {
			fmt.Println("invalid entry, likely not a string or file corruption")
			log.Fatal(errTextIn)
		}

		handle.CreateAndWriteFile(fileNameIn, fileTextIn)
	case 2:
		// add handling to create error if an existing file exists
		allTxtFiles := handle.ReadDataFolder()
		fileName, errSel := handle.Selection(allTxtFiles)
		if errSel != nil {
			fmt.Println("likely invalid selection")
			//handle this better to go back to case2
			initModules()
		}
		key, baseName := handle.EncKeyOne(fileName)

		encData := handle.EncProcTwo(fileName, key)
		handle.EncWrite(encData, baseName)

	case 3:
		fmt.Println(handle.ReadDataFolderDec())

	case 4:

		fmt.Println("4")
	case 5:
		os.Exit(3)
	default:
		fmt.Println("invalid selection")
		initModules()
	}

}

func tempReadKey() []byte {
	a, _ := os.ReadFile("ok")
	return a
}
