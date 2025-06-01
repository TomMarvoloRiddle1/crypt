package main

import (
	"crypt/internal"
	"os"
)

func main() {

	cliArg := os.Args

	mode := cliArg[1]

	switch mode {
	case "genrsa":
		internal.GenRsa()
	case "encrypt":
		fileName := cliArg[2]
		internal.Encrypt(fileName)
	case "decrypt":
		fileName := cliArg[2]
		internal.Decrypt(fileName)

	}

}
