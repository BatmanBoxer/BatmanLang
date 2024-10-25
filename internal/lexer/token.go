package lexer

import "fmt"

type TokenType = int

const (
	null = iota
	RETURN
	INT_LIT
	STRING_LIT

	FUN
	LEFT_PARAMS
	RIGHT_PARAMS
	LEFT_BRACKET
	RIGHT_BRACKET
	VAR
	PRINT
	PRINTLN
	SEMICOLON
	IDENTIFIER
	EOF

	EQUALS
	IF
	ELSE
	FOR
	WHILE

  PLUS
  MINUS
  MUNTIPLY
  DIVIDE
)

type Token struct {
	Tokentype TokenType
	Value     *string
}

func (token *Token) Debug() {
	var tokenValue string
	if token.Value != nil {
		tokenValue = *token.Value
	}

	switch token.Tokentype {
	case RETURN:
		fmt.Printf("Token Type: RETURN, Value: %s\n", tokenValue)
	case INT_LIT:
		fmt.Printf("Token Type: INT_LIT, Value: %s\n", tokenValue)
	case STRING_LIT:
		fmt.Printf("Token Type: STRING_LIT, Value: %s\n", tokenValue)
	case FUN:
		fmt.Printf("Token Type: FUN, Value: %s\n", tokenValue)
	case LEFT_PARAMS:
		fmt.Printf("Token Type: LEFT_PARAMS, Value: %s\n", tokenValue)
	case RIGHT_PARAMS:
		fmt.Printf("Token Type: RIGHT_PARAMS, Value: %s\n", tokenValue)
	case LEFT_BRACKET:
		fmt.Printf("Token Type: LEFT_BRACKET, Value: %s\n", tokenValue)
	case RIGHT_BRACKET:
		fmt.Printf("Token Type: RIGHT_BRACKET, Value: %s\n", tokenValue)
	case VAR:
		fmt.Printf("Token Type: VAR, Value: %s\n", tokenValue)
	case PRINT:
		fmt.Printf("Token Type: PRINT, Value: %s\n", tokenValue)
	case PRINTLN:
		fmt.Printf("Token Type: PRINTLN, Value: %s\n", tokenValue)
	case SEMICOLON:
		fmt.Printf("Token Type: SEMICOLON, Value: %s\n", tokenValue)
	case IDENTIFIER:
		fmt.Printf("Token Type: IDENTIFIER, Value: %s\n", tokenValue)
	case EOF:
		fmt.Println("Token Type: EOF")
	case EQUALS:
		fmt.Printf("Token Type: EQUALS, Value: %s\n", tokenValue)
	case IF:
		fmt.Printf("Token Type: IF, Value: %s\n", tokenValue)
	case ELSE:
		fmt.Printf("Token Type: ELSE, Value: %s\n", tokenValue)
	case FOR:
		fmt.Printf("Token Type: FOR, Value: %s\n", tokenValue)
	case WHILE:
		fmt.Printf("Token Type: WHILE, Value: %s\n", tokenValue)
	default:
		fmt.Printf("Token Type: UNKNOWN, Value: %s\n", tokenValue)
	}
}
