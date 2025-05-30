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

	switch mode {
	case "genrsa":
		internal.ModThree()
	case "decrypt":
		internal.ModTwo()
	case "encrypt":
		internal.ModOne()
	}

}
