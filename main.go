package main

import (
	"crypt/internal"
	"os"
)

func main() {

	cliArg := os.Args

	var inputs []string
	for _, v := range cliArg {
		inputs = append(inputs, v)
	}

	//inputs[0] always == name of ./PROGRAMNAME
	mode := inputs[1]
	fileName := inputs[2]

	switch mode {
	case "genrsa":
		internal.GenRsa()
	case "encrypt":
		internal.Encrypt(fileName)
	case "decrypt":
		internal.Decrypt(fileName)

	}

}
