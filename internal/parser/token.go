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
	Condition BinaryExpression
	Body      *Block
}

type FunctionDeclaration struct {
	Token      lexer.Token
	Name       string
	Parameters []interface{}
	Body       *Block
}
type FunctionCall struct{
  Token lexer.Token
  Parameters []interface{}
  Name string
}
type LoopStatement struct{
  Token lexer.Token
  Condition BinaryExpression
  Body *Block

}
type Expression interface{}

type VariableDeclaration struct{
  Token lexer.Token
  Name string
  Value []lexer.Token
}
type VariableReasign struct{
  Token lexer.Token
  Name string
  Value []lexer.Token
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
