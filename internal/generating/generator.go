package generating

import (
	"aurum/internal/parcing"
	"fmt"
)

type Generator struct {
	nodeProgram parcing.NodeProg
}

func NewGenerator(nodeProgram parcing.NodeProg) Generator {
	return Generator{nodeProgram}
}

func (g *Generator) Generate() string {
	s := fmt.Sprintf(
		`package main

import "os"

func main() {
%s
}

`, g.nodeProgram.Code())
	return s
}
