package mylib

import (
	"encoding/json"
	"os"
)

func WriteJson() {
	data, _ := os.ReadFile("test.json")

	type Keys struct {
		NameFile string
		Key      string
	}

	var newKeys []Keys

	_ = json.Unmarshal([]byte(data), &newKeys)

	newKeys = append(newKeys, Keys{NameFile: "yes2", Key: "key2"})

	final, _ := json.Marshal(newKeys)

	os.WriteFile("test.json", final, 0666)
}
