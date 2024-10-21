package lexer

type TokenType = int

const(
  RETURN =  iota
  INT_LIT
  SEMICOLON
  
)

type Token struct{
  tokentype TokenType;
  value *string;
}
