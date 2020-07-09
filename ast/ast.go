package ast

// TokenLiteral enables the Program struct to implement the Node interface
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// To hold the identifier of the binding, the x in let x = 5; , we have the Identifier struct type,
// which implements the Expression interface. But the identifier in a let statement doesn’t produce
// a value, right? So why is it an Expression ? It’s to keep things simple. Identifiers in other parts
// of a Monkey program do produce values, e.g.: let x = valueProducingIdentifier; . And to
// keep the number of different node types small, we’ll use Identifier here to represent the name
// in a variable binding and later reuse it, to represent an identifier as part of or as a complete
// expression.
func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
