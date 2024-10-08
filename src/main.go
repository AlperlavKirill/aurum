package main

import (
	"aurum/internal/generating"
	"aurum/internal/parcing"
	"aurum/internal/tokenizing"
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
	nodeProg := parser.ParseProg()

	generator := generating.NewGenerator(nodeProg)
	code := generator.Generate()

	writeGoFile(code)
}

func writeGoFile(code string) {
	f, err := os.Create("./output/test1.go")

	if err != nil {
		log.Fatal("Error creating a file", err)
	}

	defer f.Close()

	f.Write([]byte(code))
}

func fileContent() (string, error) {
	fileName := os.Args[1]
	contentBytes, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(contentBytes), nil
}
