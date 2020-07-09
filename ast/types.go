package ast

import "monkey-interpreter/token"

// Program is the root node of every AST our parser produces
type Program struct {
	Statements []Statement
}

// LetStatement ...
type LetStatement struct {
	Token token.Token //the token.LET token
	Name  *Identifier
	Value Expression
}

// Identifier ...
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

// ReturnStatement ...
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}
