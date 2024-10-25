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

func (parser *Parser) ParseProgram() Node {
	program := &Program{Statements: []Node{}}

	for parser.index < len(parser.tokens) {
		token := parser.peak(0)

		switch token.Tokentype {
		case lexer.FUN:
			funcDecl := parser.parseFunctionDeclaration()
			program.Statements = append(program.Statements, funcDecl)
		case lexer.EOF:
			fmt.Println("parsed sucessfully")
			parser.consume()
		default:
			//panic("Cannot have anything ouside function")
		}
	}

	return program
}

func (parser *Parser) parseFunctionDeclaration() *FunctionDeclaration {
	funToken := parser.consume()  // Consume the 'fun' keyword.
	nameToken := parser.consume() // Consume the function name.
  parser.consume()
  parser.consume()
	name := *nameToken.Value

	body := parser.parseBlock() // Parse the block of the function.

	return &FunctionDeclaration{
		Token:      funToken,
		Name:       name,
		Parameters: []string{}, // Assuming no parameters for now.
		Body:       body,
	}
}

//func (parser *Parser) parseIfStatement() *IfStatement{
//  ifToken := parser.consume()

//}

func (parser *Parser) parseBlock() *Block {
	block := &Block{Token: parser.consume(), Body: []Node{}} // Consume the '{'.
	for parser.index < len(parser.tokens) {
		token := parser.peak(0)
		// Check for the end of the block.
		if token.Tokentype == lexer.RIGHT_PARAMS {
			parser.consume() // Consume the '}'.
			break
		}

		switch token.Tokentype {
		case lexer.PRINT:
			printStmt := parser.parsePrintStatement()
			block.Body = append(block.Body, printStmt)

		case lexer.FUN:
			// Check for nested blocks.
			nestedFunction := parser.parseFunctionDeclaration() // Recursive call for nested block.
			block.Body = append(block.Body, nestedFunction)

		case lexer.IF:
			ifStatement := parser.parseIfStatement()
      block.Body = append(block.Body, ifStatement)


		case lexer.VAR:
			variableDeclaration := parser.parseVariableDeclaration()
			block.Body = append(block.Body, variableDeclaration)
		case lexer.IDENTIFIER:
			parser.consume()

		case lexer.SEMICOLON:
			parser.consume()
		default:
			fmt.Println(parser.peak(0).Tokentype)
      fmt.Println(parser.index)
			panic("Unknown Token")
		}
	}

	return block
}

func (parser *Parser) parseIfStatement() *IfStatement {
	ifToken := parser.consume()   // Consume the 'if' keyword.
	parser.consume()              // Consume the (.
	leftToken := parser.consume() //Consume the value

	parser.consume() // ==
	parser.consume() // ==
  

	rightTOken := parser.consume()
  parser.consume()  //Consume the (

  body := parser.parseBlock()
	return &IfStatement{
		Token: ifToken,
		Condition: BinaryExpression{
			Left:     leftToken,
			Right:    rightTOken,
			Operator: "==",
		},
		Body: body,
	}

}

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
func (parser *Parser) parseVariableDeclaration() *VariableDeclaration {
	varToken := parser.consume() //Consume the var keyword
	identifier := parser.consume()
	varEquals := parser.consume()  //Consume the =
	valueToken := parser.consume() //Consume the value

	if varEquals.Tokentype != lexer.EQUALS {
		varEquals.Debug()
		panic("Invalid Variable declaration")
	}

	return &VariableDeclaration{
		Token: varToken,
		Name:  *identifier.Value,
		Value: valueToken,
	}

}

func (parser *Parser) consume() lexer.Token {
	token := parser.tokens[parser.index]
	parser.index++
	return token
}

func (parser *Parser) peak(i int) lexer.Token {
	return parser.tokens[parser.index+i]
}
