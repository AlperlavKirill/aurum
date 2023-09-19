package generating

import (
	"aurum/internal/parcing"
	"fmt"
)

type Generator struct {
	nodeQuit parcing.NodeQuit
}

func NewGenerator(nodeQuit parcing.NodeQuit) Generator {
	return Generator{nodeQuit}
}

func (g *Generator) Generate() string {
	s := fmt.Sprintf(
		`
package main

import "os"

func main() {
	%s
}

`, g.nodeQuit.Code())
	return s
}
