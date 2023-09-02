package src

import (
	tokenizer2 "aurum/internal/tokenizer"
	"fmt"
	"os"
)

func main() {
	fileName := os.Args[1]
	contentBytes, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Errorf("error reading file %s", fileName)
		os.Exit(-1)
	}

	contents := string(contentBytes)

	_ = tokenizer2.NewTokenizer(contents)
}
