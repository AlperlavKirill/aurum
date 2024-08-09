package parcing

import (
	"aurum/internal/tokenizing"
	"errors"
)

type Parser struct {
	tokens []tokenizing.Token
	index  int
}

func NewParser(tokens []tokenizing.Token) Parser {
	return Parser{tokens, 0}
}

func (p *Parser) parseTerm() (NodeTerm, error) {
	var nodeTerm NodeTerm
	if t, err := p.tryConsume(tokenizing.Ident); err == nil {
		var identTerm = NodeTermIdent{ident: t}
		nodeTerm.term = &identTerm
	} else if t, err = p.tryConsume(tokenizing.IntLit); err == nil {
		var intLitTerm = NodeTermIntLit{intLit: t}
		nodeTerm.term = &intLitTerm
	} else {
		return nodeTerm, errors.New("error in parseTerm()")
	}
	return nodeTerm, nil
}

func (p *Parser) parseExpr() (NodeExpr, error) {
	if nt, err := p.parseTerm(); err == nil {
		return NodeExpr{expr: &nt}, nil
	}
	return NodeExpr{}, errors.New("error in parseExpr()")
}

func (p *Parser) parseStmt() (NodeStmt, error) {
	var nodeStmt NodeStmt
	if p.isOffsetValid(0, tokenizing.Let) &&
		p.isOffsetValid(1, tokenizing.Ident) &&
		p.isOffsetValid(2, tokenizing.Eq) {
		p.consume()             // consume let
		letToken := p.consume() // consume identifier
		p.consume()             // consume =
		if ne, err := p.parseExpr(); err == nil {
			nsl := NodeStmtLet{ident: letToken, expr: ne}
			nodeStmt.stmt = &nsl
		}
		p.consume() // consume ;
		return nodeStmt, nil
	}
	if p.isOffsetValid(0, tokenizing.Quit) {
		p.consume() // consume quit
		if ne, err := p.parseExpr(); err == nil {
			nsq := NodeStmtQuit{expr: ne}
			nodeStmt.stmt = &nsq
		}
		return nodeStmt, nil
	}
	return NodeStmt{}, errors.New("invalid statement")
}

func (p *Parser) ParseProg() NodeProg {
	var nodeProg = NewNodeProg()
	for _, err := p.peak(); err == nil; _, err = p.peak() {
		ns, _ := p.parseStmt()
		nodeProg.pushBack(&ns)
	}
	return nodeProg
}

func (p *Parser) peak() (tokenizing.Token, error) {
	return p.peakOffset(0)
}

func (p *Parser) peakOffset(offset int) (tokenizing.Token, error) {
	if p.index+offset >= len(p.tokens)-1 {
		return tokenizing.Token{}, errors.New("ERROR: offset out of bounds parsing tokens")
	}
	return p.tokens[p.index+offset], nil
}

func (p *Parser) isOffsetValid(offset int, tokenType tokenizing.TokenType) bool {
	token, err := p.peakOffset(offset)
	if err != nil {
		return false
	}
	if token.Type == tokenType {
		return true
	}
	return false
}

func (p *Parser) tryConsume(tokenType tokenizing.TokenType) (tokenizing.Token, error) {
	if t, err := p.peak(); err == nil && t.Type == tokenType {
		return p.consume(), nil
	}
	return tokenizing.Token{}, errors.New("error in try_consume")
}

func (p *Parser) consume() tokenizing.Token {
	token := p.tokens[p.index]
	p.index++
	return token
}
