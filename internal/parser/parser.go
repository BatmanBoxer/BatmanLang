package parser

import "compileringo/internal/lexer"

type Parser struct {
	tokens []lexer.Token
	index  int
}

func NewParser(tokens []lexer.Token) *Parser {
	return &Parser{
		tokens: tokens,
		index:  0,
	}
}

func (parser *Parser) ParseProgram() Node {
	program := &Program{Statements: []Node{}}

	for parser.index < len(parser.tokens) {
		token := parser.tokens[parser.index]

		switch token.Tokentype {
		case lexer.FUN:
			funcDecl := parser.parseFunctionDeclaration()
			program.Statements = append(program.Statements, funcDecl)

		default:
			parser.index++
		}
	}

	return program
}

func (parser *Parser) parseFunctionDeclaration() *FunctionDeclaration {
	funToken := parser.tokens[parser.index]
	parser.index++

	nameToken := parser.tokens[parser.index]
	name := *nameToken.Value
	parser.index++

	body := parser.parseBlock()

	return &FunctionDeclaration{
		Token:      funToken,
		Name:       name,
		Parameters: []string{},
		Body:       body,
	}
}

func (parser *Parser) parseBlock() *Block {
  block := &Block{Token:parser.tokens[parser.index], Body: []Node{}}

	parser.index++

	for parser.index < len(parser.tokens) {
		token := parser.tokens[parser.index]
		//needs to be change to supports nesting and to lazy to do that now so i dont care
		if token.Tokentype == lexer.RIGHT_BRACKET {
			break
		}

		switch token.Tokentype {
		case lexer.PRINT:
			printStmt := parser.parsePrintStatement()
			block.Body = append(block.Body, printStmt)

		default:
			parser.index++
		}
	}

	parser.index++
	return block
}

func (parser *Parser) parsePrintStatement() *PrintStatement {
	printToken := parser.tokens[parser.index]
	parser.index++

	valueToken := parser.tokens[parser.index]
	value := &Literal{
		Token: valueToken,
		Value: *valueToken.Value,
	}

	parser.index++

	return &PrintStatement{
		Token: printToken,
		Value: value,
	}
}
