package tokenizing

type TokenType int8

const (
	Quit TokenType = iota
	IntLit
	Semi
)

type Token struct {
	Type  TokenType
	Value *string
}

func (token *Token) HasValue() bool {
	return token.Value != nil
}