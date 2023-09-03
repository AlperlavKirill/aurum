package main

import (
	"aurum/internal/parcing"
	"aurum/internal/tokenizing"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Incorrect usage, pass <file.au> filename")
	}

	content, err := fileContent()

	if err != nil {
		log.Fatal("Error reading file")
	}

	tokenizer := tokenizing.NewTokenizer(content)
	tokens := tokenizer.Tokenize()

	parser := parcing.NewParser(tokens)
	nodeQuit := parser.Parse()

	fmt.Printf("%+v\n", nodeQuit)
}

func fileContent() (string, error) {
	fileName := os.Args[1]
	contentBytes, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(contentBytes), nil
}
