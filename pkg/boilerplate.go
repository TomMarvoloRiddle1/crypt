package pkg

import (
	"fmt"
)

func NameInput(mode string) (string, error) {
	//this function can be replaced for os.Args later on anyways

	whatName := fmt.Sprintf("enter name of file to %s [without .txt]", mode)
	fmt.Println(whatName)

	var namePlainTxt string
	fmt.Scan(&namePlainTxt)

	return namePlainTxt, nil

}
