package ast

import (
	"github.com/komuro-hiraku/monkey/token"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// AST の ルートノード
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// Let ノード
type LetStatement struct {
	Token token.Token // token.LET トークン
	Name  *Identifier
	Value Expression
}

// Statement, Node インタフェースを満たすための実装
func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Return ノード
type ReturnStatement struct {
	Token       token.Token // 'return' トークン
	ReturnValue Expression
}

// Statement, Node インタフェースを満たすための実装
func (rt *ReturnStatement) statementNode()       {}
func (rt *ReturnStatement) TokenLiteral() string { return rt.Token.Literal }

type Identifier struct {
	Token token.Token // token.IDENT トークン
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
