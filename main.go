package main

import (
	"crypt/pkg"
	"fmt"
	"os"
)

func main() {
	initModules(initSelect())
}

func initSelect() int {

	selectionOpt := `1) Create and write to text file
2) Encrypt text file
3) Decrypt text file
4) Exit program`
	fmt.Println(selectionOpt)

	var selectionUser int
	fmt.Scan(&selectionUser)

	return selectionUser
}

func initModules(selectionUser int) {

	switch selectionUser {
	case 1:
		pkg.CreateAndWriteFile(pkg.PromptData())
	case 2:
		fileName, errSel := pkg.Selection()
		//issue here, panic: runtime error: index out of range [7] with length 3
		if errSel != nil {
			fmt.Println("likely invalid selection")
			//add error handling, to return to case2
			initModules(2)
		}

		key, baseName := pkg.EncKeyOne(fileName)

		encData := pkg.EncProcTwo(fileName, key)
		pkg.EncWrite(encData, baseName)

	case 3:
		pkg.DecFlow()

	case 4:
		os.Exit(3)

	default:
		fmt.Println("invalid selection")
		initModules(initSelect())
	}

}
