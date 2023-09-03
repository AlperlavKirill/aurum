package parcing

import (
	"aurum/internal/tokenizing"
	"errors"
	"log"
)

type Parser struct {
	tokens []tokenizing.Token
	index  int
}

func NewParser(tokens []tokenizing.Token) Parser {
	return Parser{tokens, 0}
}

func (p *Parser) parseExpr() (NodeExpr, error) {
	if token, err := p.peak(); err == nil && token.Type == tokenizing.IntLit && token.HasValue() {
		return NodeExpr{intLit: p.consume()}, nil
	} else {
		return NodeExpr{}, errors.New("ERROR: error parsing expression node")
	}
}

func (p *Parser) Parse() NodeQuit {
	var nodeQuit NodeQuit
	for token, err := p.peak(); err == nil; token, err = p.peak() {
		if token.Type == tokenizing.Quit {
			p.consume()
			if nodeExpr, err := p.parseExpr(); err == nil {
				nodeQuit.expr = nodeExpr
			} else {
				log.Fatal("ERROR: error parsing program")
			}
		} else if token.Type == tokenizing.Semi {
			p.consume()
		} else {
			log.Fatal("ERROR: incorrect usage")
		}
	}
	return nodeQuit
}

func (p *Parser) peak() (tokenizing.Token, error) {
	return p.peakOffset(0)
}

func (p *Parser) peakOffset(offset int) (tokenizing.Token, error) {
	if p.index+offset >= len(p.tokens) {
		return tokenizing.Token{}, errors.New("ERROR: offset out of bounds parsing tokens")
	}
	return p.tokens[p.index+offset], nil
}

func (p *Parser) consume() tokenizing.Token {
	token := p.tokens[p.index]
	p.index++
	return token
}
