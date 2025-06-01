package main

import (
	"crypt/internal/appFlow"
	"fmt"
	"os"
)

func main() {

	cliArg := os.Args

	mode := cliArg[1]

	switch mode {
	case "help":
		fmt.Println("gonna add this bitch later")
	case "genrsa":
		appFlow.GenRsa()
	case "encrypt":
		fileName := cliArg[2]
		appFlow.Encrypt(fileName)
	case "decrypt":
		fileName := cliArg[2]
		appFlow.Decrypt(fileName)
	case "zip":
		folderName := cliArg[2]
		appFlow.Zip(folderName)

	default:
		fmt.Println("invalid selection")
	}

}
