package parser

import (
	"github.com/komuro-hiraku/monkey/ast"
	"github.com/komuro-hiraku/monkey/lexer"
	"github.com/komuro-hiraku/monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser {l: l}

	// 2つトークンを読み込む
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
