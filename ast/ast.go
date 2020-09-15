package ast

import (
	"bytes"
	"github.com/komuro-hiraku/monkey/token"
)

type Node interface {
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	var out bytes.Buffer

	// Statements の中身をBufferに詰めて文字列にする
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// Let ノード
type LetStatement struct {
	Token token.Token // token.LET トークン
	Name  *Identifier
	Value Expression
}

// Let: Statement, Node インタフェースを満たすための実装
func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

// Return ノード
type ReturnStatement struct {
	Token       token.Token // 'return' トークン
	ReturnValue Expression
}

// Return: Statement, Node インタフェースを満たすための実装
func (rt *ReturnStatement) statementNode()       {}
func (rt *ReturnStatement) TokenLiteral() string { return rt.Token.Literal }
func (rt *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rt.TokenLiteral() + " ")

	if rt.ReturnValue != nil {
		out.WriteString(rt.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// Identifier ノード
type Identifier struct {
	Token token.Token // token.IDENT トークン
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string {
	return i.Value
}

// Expression ノード
type ExpressionStatement struct {
	Token token.Token
	Expression Expression
}

func (ex *ExpressionStatement) expressionNode() {}
func (ex *ExpressionStatement) TokenLiteral() string { return ex.Token.Literal }
func (ex *ExpressionStatement) String() string {
	if ex.Expression != nil {
		return ex.Expression.String()
	}
	return ""
}