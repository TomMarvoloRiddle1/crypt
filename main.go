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
		fileName, _ := pkg.Selection()

		pkg.EntireEnc(fileName)

	case 3:
		//broken af
		pkg.DecText()

	case 4:
		os.Exit(3)

	default:
		fmt.Println("invalid selection")
		initModules(initSelect())
	}

}
