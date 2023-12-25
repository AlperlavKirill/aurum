package tokenizing

import (
	"errors"
	"log"
	"unicode"
)

type Tokenizer struct {
	contents []rune
	index    int
}

func NewTokenizer(contents string) Tokenizer {
	return Tokenizer{[]rune(contents), 0}
}

func (t *Tokenizer) peekOffset(offset int) (rune, error) {
	if t.index+offset >= len(t.contents) {
		return 'f', errors.New("ERROR: peek out of range")
	}
	return t.contents[t.index+offset], nil
}

func (t *Tokenizer) peek() (rune, error) {
	return t.peekOffset(0)
}

func (t *Tokenizer) consume() rune {
	char := t.contents[t.index]
	t.index++
	return char
}

func (t *Tokenizer) Tokenize() []Token {
	var tokens []Token
	var buf []rune

	clearBuf := func() {
		buf = []rune{}
	}

	consumeToBuf := func() {
		buf = append(buf, t.consume())
	}

	isAlpha := func(r rune) bool {
		return unicode.IsLetter(r)
	}

	isAlphaNum := func(r rune) bool {
		return unicode.IsDigit(r) || unicode.IsLetter(r)
	}

	isDigit := func(r rune) bool {
		return unicode.IsDigit(r)
	}

	for char, err := t.peek(); err == nil; char, err = t.peek() {
		if isAlpha(char) {
			consumeToBuf()
			for char, err = t.peek(); err == nil && isAlphaNum(char); char, err = t.peek() {
				consumeToBuf()
			}
			if string(buf) == "quit" {
				tokens = append(tokens, Token{Type: Quit, Value: nil})
				clearBuf()
				continue
			}
			if string(buf) == "let" {
				tokens = append(tokens, Token{Type: Let, Value: nil})
				clearBuf()
				continue
			} else {
				identName := string(buf)
				tokens = append(tokens, Token{Type: Ident, Value: &identName})
				continue
			}
		} else if isDigit(char) {
			consumeToBuf()
			for char, err = t.peek(); err == nil && isDigit(char); char, err = t.peek() {
				consumeToBuf()
			}
			intLitValue := string(buf)
			tokens = append(tokens, Token{IntLit, &intLitValue})
			clearBuf()
			continue
		} else if char == '=' {
			t.consume()
			tokens = append(tokens, Token{Eq, nil})
			continue
		} else if char == ';' {
			t.consume()
			tokens = append(tokens, Token{Semi, nil})
			continue
		} else if char == ' ' {
			t.consume()
			continue
		} else {
			log.Fatal("SYNTAX ERROR: you messed up")
		}
	}
	t.index = 0
	return tokens
}
