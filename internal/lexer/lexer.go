package lexer

import (
	"errors"
	"fmt"
	"unicode"
)

type Lexer struct {
	src   string
	buf   []rune
	index int
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		src: input,
		buf: []rune{},
	}
}

func (lexer *Lexer) Tokenize() []Token {

	var tokens []Token
	for {
		token, err := lexer.peak()
		if err != nil {
			fmt.Print(err)
			break
		}
		if unicode.IsLetter(rune(token)) {
			tokens = append(tokens, lexer.lexAlpha())
		} else if unicode.IsDigit(rune(token)) {
			tokens = append(tokens, lexer.lexDigit())
		} else if token == '"' {
			tokens = append(tokens, lexer.lexString())
		} else if token == ' ' {
			lexer.consume()
		} else if rune(token) == '(' {
			lexer.consume()
			tokens = append(tokens, Token{LEFT_BRACKET, nil})
		} else if rune(token) == ')' {
			lexer.consume()
			tokens = append(tokens, Token{RIGHT_BRACKET, nil})
		} else if rune(token) == '{' {
			lexer.consume()
			tokens = append(tokens, Token{LEFT_PARAMS, nil})
		} else if rune(token) == '}' {
			lexer.consume()
			tokens = append(tokens, Token{RIGHT_PARAMS, nil})
		} else if rune(token) == ';' {
			lexer.consume()
			tokens = append(tokens, Token{SEMICOLON, nil})
		} else if rune(token) == '+' {
			lexer.consume()
			tokens = append(tokens, Token{PLUS, nil})

		} else if rune(token) == '-' {
			lexer.consume()
			tokens = append(tokens, Token{MINUS, nil})

		} else if rune(token) == '*' {
			lexer.consume()
			tokens = append(tokens, Token{MULTIPLY, nil})

		} else if rune(token) == '/' {
			lexer.consume()
			tokens = append(tokens, Token{DIVIDE, nil})

		} else if rune(token) == '=' {
			lexer.consume()
			tokens = append(tokens, Token{EQUALS, nil})
		} else if rune(token) == '\n' {
			if lexer.checkEOF() {
				tokens = append(tokens, Token{EOF, nil})
				break
			} else {
				lexer.consume()
			}
		} else {
			break
		}
	}
	return tokens
}
func (lexer *Lexer) checkEOF() bool {
	return lexer.index == len(lexer.src)-1
}
func (lexer *Lexer) lexAlpha() Token {
	parsertoken := Token{-1, nil}
	lexer.buf = append(lexer.buf, rune(lexer.consume()))
	for {
		token, err := lexer.peak()
		if err != nil {
			break
		}
		if unicode.IsLetter(rune(token)) || unicode.IsDigit(rune(token)) {
			lexer.buf = append(lexer.buf, rune(lexer.consume()))
		} else {
			break
		}
	}

	switch string(lexer.buf) {
	case "return":
		parsertoken = Token{RETURN, nil}
		lexer.buf = lexer.buf[:0]

	case "fun":
		parsertoken = Token{FUN, nil}
		lexer.buf = lexer.buf[:0]

	case "var":
		parsertoken = Token{VAR, nil}
		lexer.buf = lexer.buf[:0]

	case "print":
		parsertoken = Token{PRINT, nil}
		lexer.buf = lexer.buf[:0]

	case "println":
		parsertoken = Token{PRINTLN, nil}
		lexer.buf = lexer.buf[:0]

	case "if":
		parsertoken = Token{IF, nil}
		lexer.buf = lexer.buf[:0]

	case "else":
		parsertoken = Token{ELSE, nil}
		lexer.buf = lexer.buf[:0]

	case "for":
		parsertoken = Token{FOR, nil}
		lexer.buf = lexer.buf[:0]

	case "while":
		parsertoken = Token{WHILE, nil}
		lexer.buf = lexer.buf[:0]

	default:
		var strbuf = string(lexer.buf)
		parsertoken = Token{IDENTIFIER, &strbuf}
		lexer.buf = lexer.buf[:0]
	}
	return parsertoken
}
func (lexer *Lexer) lexString() Token {
	parsedtoken := Token{-1, nil}
	lexer.buf = append(lexer.buf, rune(lexer.consume()))
	for {
		token, err := lexer.peak()
		if err != nil {
			break
		}
		if rune(token) != '"' {
			lexer.buf = append(lexer.buf, rune(lexer.consume()))
		} else {
			lexer.buf = append(lexer.buf, rune(lexer.consume()))
			var temp = string(lexer.buf)
			parsedtoken = Token{STRING_LIT, &temp}
			lexer.buf = lexer.buf[:0]
			break
		}
	}
	return parsedtoken
}

func (lexer *Lexer) lexDigit() Token {
	parsertoken := Token{-1, nil}
	for {
		token, err := lexer.peak()
		if err != nil {
			break
		}
		if !unicode.IsDigit(rune(token)) {
			value := string(lexer.buf)
			parsertoken = Token{INT_LIT, &value}
			lexer.buf = lexer.buf[:0]
			break

		} else {
			lexer.buf = append(lexer.buf, rune(lexer.consume()))
		}
	}
	return parsertoken
}
func (lexer *Lexer) peak() (byte, error) {
	var err = errors.New("Failed lexing")
	var character byte
	if lexer.index > len(lexer.src) {
		return character, err
	} else {
		return lexer.src[lexer.index], nil
	}
}

func (lexer *Lexer) consume() byte {
	character := lexer.src[lexer.index]
	lexer.index++
	return character
}
