package parser

import (
	"compileringo/internal/lexer"
	"fmt"
)

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

func (parser *Parser) consume() lexer.Token {
	token := parser.tokens[parser.index]
	parser.index++
	return token
}

func (parser *Parser) peak() lexer.Token {
	return parser.tokens[parser.index]
}

func (parser *Parser) ParseProgram() Node {
	program := &Program{Statements: []Node{}}

	for parser.index < len(parser.tokens) {
		token := parser.peak()

		switch token.Tokentype {
		case lexer.FUN:
			funcDecl := parser.parseFunctionDeclaration()
			program.Statements = append(program.Statements, funcDecl)

		default:
			parser.consume()
		}
	}

	return program
}

func (parser *Parser) parseFunctionDeclaration() *FunctionDeclaration {
	funToken := parser.consume()  // Consume the 'fun' keyword.
	nameToken := parser.consume() // Consume the function name.
	name := *nameToken.Value

	body := parser.parseBlock() // Parse the block of the function.

	return &FunctionDeclaration{
		Token:      funToken,
		Name:       name,
		Parameters: []string{}, // Assuming no parameters for now.
		Body:       body,
	}
}

// parseBlock parses a block of code and returns it as a Block node.
func (parser *Parser) parseBlock() *Block {
	block := &Block{Token: parser.consume(), Body: []Node{}} // Consume the '{'.
	for parser.index < len(parser.tokens) {
		token := parser.peak() // Use peak to check the next token.

		// Check for the end of the block.
		if token.Tokentype == lexer.RIGHT_PARAMS {
			parser.consume() // Consume the '}'.
			fmt.Println("reached here")
			break
		}

		switch token.Tokentype {
		case lexer.PRINT:
			printStmt := parser.parsePrintStatement()
			block.Body = append(block.Body, printStmt)

		default:
			parser.consume() // Skip unknown token.
		}
	}

	return block
}

// parsePrintStatement parses a print statement and returns it.
func (parser *Parser) parsePrintStatement() *PrintStatement {
	printToken := parser.consume()       // Consume the 'print' keyword.
	leftbraketToken := parser.consume()  // Consume the (.
	valueToken := parser.consume()       //Consume the value
	rightbraketToken := parser.consume() //Consume the ).

	if leftbraketToken.Tokentype != lexer.LEFT_BRACKET && rightbraketToken.Tokentype != lexer.RIGHT_BRACKET {
		panic("invalid use of print")
	}

	value := &Literal{
		Token: valueToken,
		Value: *valueToken.Value,
	}

	return &PrintStatement{
		Token: printToken,
		Value: value,
	}
}
