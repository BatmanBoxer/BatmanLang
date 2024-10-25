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
			panic("Cannot have anything ouside function")
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
		Token: funToken,
		Name:  name,
		// Assuming no parameters for now.
		Body: body,
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

		case lexer.WHILE:
			loopStatement := parser.parseLoopStatement()
			block.Body = append(block.Body, loopStatement)
		case lexer.VAR:
			variableDeclaration := parser.parseVariableDeclaration()
			block.Body = append(block.Body, variableDeclaration)
		case lexer.IDENTIFIER:
			if parser.isfunCall() {
				functionCall := parser.parseFunCall()
				block.Body = append(block.Body, functionCall)
			}
			if parser.isVariableReasign() {
				variableReasign := parser.parseVariableReasign()
				block.Body = append(block.Body, variableReasign)
			}
		case lexer.SEMICOLON:
			parser.consume()
		default:
			panic("unknown token here")
		}
	}

	return block
}
func (parser *Parser) isfunCall() bool {
	if parser.peak(1).Tokentype == lexer.LEFT_BRACKET {
		return true
	} else {
		return false
	}
}
func (parser *Parser) isVariableReasign() bool {
	if parser.peak(1).Tokentype == lexer.EQUALS {
		return true
	} else {
		return false
	}
}

func (parser *Parser) parseIfStatement() *IfStatement {
	ifToken := parser.consume()   // Consume the 'if' keyword.
	parser.consume()              // Consume the (.
	leftToken := parser.consume() //Consume the value

	parser.consume() // ==
	parser.consume() // ==

	rightTOken := parser.consume()
	parser.consume() //Consume the (

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

func (parser *Parser) parseFunCall() *FunctionCall {
	functionToken := parser.consume()
	parser.consume()
	parser.consume()

	return &FunctionCall{
		Token: functionToken,
		Name:  *functionToken.Value,
	}
}
func (parser *Parser) parseLoopStatement() *LoopStatement {
	loopToken := parser.consume() // Consume the 'if' keyword.
	parser.consume()              // Consume the (.
	leftToken := parser.consume() //Consume the value

	parser.consume() // ==
	parser.consume() // ==

	rightTOken := parser.consume()
	parser.consume() //Consume the (

	body := parser.parseBlock()
	return &LoopStatement{
		Token: loopToken,
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
	varEquals := parser.consume() //Consume the =
	expression := []lexer.Token{}
	for {
		if parser.peak(0).Tokentype == lexer.SEMICOLON {
			break
		}
		expression = append(expression, parser.consume())
	}

	if varEquals.Tokentype != lexer.EQUALS {
		panic("Invalid Variable declaration")
	}
	return &VariableDeclaration{
		Token: varToken,
		Name:  *identifier.Value,
		Value: expression,
	}
}
func (parser *Parser) parseVariableReasign() *VariableReasign {
	varName := parser.consume() // consume name
	parser.consume()            //consume =
	expression := []lexer.Token{}
	for {
		if parser.peak(0).Tokentype == lexer.SEMICOLON {
			break
		}
		expression = append(expression, parser.consume())
	}
	return &VariableReasign{
		Token: varName,
		Name:  *varName.Value,
		Value: expression,
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
