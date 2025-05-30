package main

import (
	"crypt/internal"
	"fmt"
	"os"
)

func main() {
	initModules(initSelect())
}

func initSelect() int {
	selectionOpt := `
1) Encrypt text file
2) Decrypt text file
3) Create RSA keypair
4) Exit program`
	fmt.Println(selectionOpt)
	var selectionUser int
	fmt.Scan(&selectionUser)
	return selectionUser
}

func initModules(selectionUser int) {

	switch selectionUser {
	case 1:
		internal.ModOne()

	case 2:

		internal.ModTwo()

	case 3:
		internal.ModThree()

	case 4:
		os.Exit(3)

	default:
		fmt.Println("invalid selection")
		initModules(initSelect())
	}

}
