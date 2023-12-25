package parcing

import "aurum/internal/tokenizing"

type NodeExpr struct {
	intLit tokenizing.Token
}

type NodeInitVar struct {
}

type NodeQuit struct {
	expr NodeExpr
}

func (nq *NodeQuit) Code() string {
	return "os.Exit(" + *nq.expr.intLit.Value + ")"
}

type CompiledNode interface {
	Code() string
}
