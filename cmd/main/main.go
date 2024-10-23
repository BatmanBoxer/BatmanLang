package main

import (
	"compileringo/internal/lexer"
	"compileringo/internal/parser"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("you messed up")
		os.Exit(1)
	} else {
		file, err := os.ReadFile(os.Args[1])
		filestring := string(file)

		tokenizer := lexer.NewLexer(filestring)
		tokens := tokenizer.Tokenize()
  
    parser := parser.NewParser(tokens)
    ProgramNode := parser.ParseProgram()

    fmt.Println(ProgramNode)
    //for _,token:= range tokens{
		 //  token.Debug()  
		//}
		if err != nil {
			fmt.Println("file path invalid")
			os.Exit(1)
		}
	}

}
