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

// ExpressionStatement ...
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

// IntegerLiteral ...
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

type PrefixExpression struct {
	Token    token.Token // The prefix token, e.g. !
	Operator string
	Right    Expression
}

type InfixExpression struct {
	Token    token.Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}
