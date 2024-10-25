package interpreter

import (
	"compileringo/internal/lexer"
	"compileringo/internal/parser"
	"fmt"
	"reflect"
)

// Interpreter struct to manage function and variable maps.
type Interpreter struct {
	FunctionMap map[string]*parser.FunctionDeclaration
	VariableMap map[string]interface{}
}

// NewInterpreter creates a new instance of the interpreter.
func NewInterpreter() *Interpreter {
	return &Interpreter{
		FunctionMap: make(map[string]*parser.FunctionDeclaration),
		VariableMap: make(map[string]interface{}),
	}
}

// VisitNode traverses the AST and executes or stores functions/variables.
func (interpreter *Interpreter) VisitNode(node parser.Node) {
	switch n := node.(type) {

	case *parser.Program:
		for _, stmt := range n.Statements {
			interpreter.VisitNode(stmt)
		}
		if mainFunc, exists := interpreter.FunctionMap["main"]; exists {
			interpreter.VisitNode(mainFunc.Body)
		} else {
			fmt.Println("No Main Function ")
		}

	case *parser.FunctionDeclaration:
		interpreter.FunctionMap[n.Name] = n

	case *parser.FunctionCall:
		if fn, exists := interpreter.FunctionMap[n.Name]; exists {
			interpreter.VisitNode(fn.Body)
		} else {
		}

	case *parser.VariableDeclaration:
		if _, exists := interpreter.VariableMap[n.Name]; exists {
			panic("cannot declare a variable twice")
		}
		interpreter.VariableMap[n.Name] = n.Value

	case *parser.VariableReasign:
		if _, exists := interpreter.VariableMap[n.Name]; exists {
			interpreter.VariableMap[n.Name] = n.Value
		} else {
			panic("variable reasinged without initializing")
		}

	case *parser.PrintStatement:
		fmt.Println("Printing: ")
		n.Token.Debug()

	case *parser.Block:
		for _, block := range n.Body {
			interpreter.VisitNode(block)
		}
	case *parser.IfStatement:
		var leftcondition interface{}
		var rightcondition interface{}

		switch value := n.Condition.Left.(type) {
		case lexer.Token:
			if value.Tokentype == lexer.IDENTIFIER {
				fmt.Println("identifier")
			}
      leftcondition = *value.Value
		}
		switch value := n.Condition.Right.(type) {
		case lexer.Token:
			fmt.Println(*value.Value)
			if value.Tokentype == lexer.IDENTIFIER {
				fmt.Println("identifier")
			}
      rightcondition = *value.Value
		}

    if leftcondition == rightcondition {
      interpreter.VisitNode(n.Body)
    }

		fmt.Println("printing val")

		if n.Condition.Left == n.Condition.Right {
			interpreter.VisitNode(n.Body)
		}

	default:
		fmt.Printf("Unknown node type: %s\n", reflect.TypeOf(n).String())
	}
}
