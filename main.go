package main

import (
	"crypt/pkg"
	"fmt"
	"os"
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
		pkg.CreateAndWriteFile(pkg.PromptData())
	case 2:
		// add handling to create error if an existing file exists
		fileName, errSel := pkg.Selection()
		if errSel != nil {
			fmt.Println("likely invalid selection")
			//add error handling, to return to case2
			initModules()
		}

		key, baseName := pkg.EncKeyOne(fileName)

		encData := pkg.EncProcTwo(fileName, key)
		pkg.EncWrite(encData, baseName)

	case 3:
		fmt.Println("3")

	case 4:
		os.Exit(3)

	default:
		fmt.Println("invalid selection")
		initModules()
	}

}
