package parcing

import (
	"aurum/internal/tokenizing"
	"fmt"
	"strings"
)

// NodeTerm
// term is either identifier
// or number
type NodeTerm struct {
	term nodeTerm
}

type NodeExpr struct {
	// expr is NodeTerm
	expr nodeExpr
}

// NodeStmt
// statement is one of these constructions:
// quit...
// let...
type NodeStmt struct {
	stmt nodeStmt
}

type NodeTermIntLit struct {
	intLit tokenizing.Token
}

type NodeTermIdent struct {
	ident tokenizing.Token
}

type NodeStmtLet struct {
	ident tokenizing.Token
	expr  NodeExpr
}

type NodeStmtQuit struct {
	expr NodeExpr
}

type NodeProg struct {
	stmts []*NodeStmt
}

/*
*
interface to define expression nodes
like NodeTerm or NodeBinExpr(coming soon)
*/
type nodeExpr interface {
	codeExpr() string
}

func (nt *NodeTerm) codeExpr() string {
	return nt.term.codeTerm()
}

/*
*
interface to define term nodes
like NodeTermIntLit, NodeTermIdent
*/
type nodeTerm interface {
	codeTerm() string
}

func (nti *NodeTermIdent) codeTerm() string {
	return *nti.ident.Value
}

func (ntil *NodeTermIntLit) codeTerm() string {
	return *ntil.intLit.Value
}

/*
*
interface to define statement nodes
like NodeStmtQuit, NodeStmtLet
*/
type nodeStmt interface {
	codeStmt() string
}

func (nsq *NodeStmtQuit) codeStmt() string {
	return fmt.Sprintf("\tos.Exit(%s)", nsq.expr.expr.codeExpr())
}

func (nsl *NodeStmtLet) codeStmt() string {
	return fmt.Sprintf("\t%s := %s", *nsl.ident.Value, nsl.expr.expr.codeExpr())
}

func NewNodeProg() NodeProg {
	return NodeProg{
		stmts: make([]*NodeStmt, 0),
	}
}

func (np *NodeProg) Code() string {
	statementCodes := make([]string, 0)
	for _, stmt := range np.stmts {
		statementCodes = append(statementCodes, stmt.stmt.codeStmt())
	}
	return strings.Join(statementCodes, "\n")
}

func (np *NodeProg) pushBack(stmt *NodeStmt) {
	np.stmts = append(np.stmts, stmt)
}
