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

var lexer Lexer = Lexer{}

func (lexer *Lexer) Tokenize() []Token {
	tokens := []Token{}
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
		} else if rune(token) == ' ' {
			lexer.consume()
		} else if rune(token) == ';' {
			lexer.consume()
			tokens = append(tokens, Token{SEMICOLON, nil})
		} else {
			break
		}
	}
	return tokens
}
func (lexer *Lexer) lexAlpha() Token {
	parsertoken := Token{}
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
		if string(lexer.buf) == "return" {
			parsertoken = Token{RETURN, nil}
		}
	}
	return parsertoken
}
func (lexer *Lexer) lexDigit() Token {
	parsertoken := Token{}
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
