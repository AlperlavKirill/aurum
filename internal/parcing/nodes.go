package parcing

import "aurum/internal/tokenizing"

type NodeExpr struct {
	intLit tokenizing.Token
}

type NodeQuit struct {
	expr NodeExpr
}
