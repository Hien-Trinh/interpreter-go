package parser

import (
	"github.com/Hien-Trinh/interpreter-go/ast"
	"github.com/Hien-Trinh/interpreter-go/lexer"
	"github.com/Hien-Trinh/interpreter-go/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

// New creates a new Parser instance
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()
	return p
}

// nextToken advances the tokens in the parser
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram parses the program
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
