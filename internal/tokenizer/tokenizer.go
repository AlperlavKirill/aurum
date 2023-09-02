package tokenizer

type TokenType int8

const (
	quit TokenType = iota
	intLit
	semi
)

type Token struct {
	Type  TokenType
	Value string
}

type Tokenizer struct {
	contents string
	index    int8
}

func NewTokenizer(contents string) Tokenizer {
	return Tokenizer{contents: contents, index: 0}
}
