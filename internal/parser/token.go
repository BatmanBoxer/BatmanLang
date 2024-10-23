package parser

import "compileringo/internal/lexer"

type Node interface{}

type Program struct {
	Statements []Node
}

type Block struct {
	Token lexer.Token
	Body  []Node
}

type Assignment struct {
	Token      lexer.Token
	Identifier string
	Value      Node
}

type IfStatement struct {
	Token     lexer.Token
	Condition Node
	Body      []Node
}

type FunctionDeclaration struct {
	Token      lexer.Token
	Name       string
	Parameters []string
	Body       *Block
}

type ReturnStatement struct {
	Token lexer.Token
	Value Node
}

type BinaryExpression struct {
	Token    lexer.Token
	Left     Node
	Operator string
	Right    Node
}

type Literal struct {
	Token lexer.Token
	Value interface{}
}

type Identifier struct {
	Token lexer.Token
	Value string
}
type PrintStatement struct {
	Token lexer.Token
	Value Node
}
