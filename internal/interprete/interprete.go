package interpreter

import (
	"compileringo/internal/lexer"
	"compileringo/internal/parser"
	"fmt"
	"reflect"
	"strings"
)

type Interpreter struct {
	FunctionMap map[string]*parser.FunctionDeclaration
	VariableMap map[string]interface{}
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		FunctionMap: make(map[string]*parser.FunctionDeclaration),
		VariableMap: make(map[string]interface{}),
	}
}

func (interpreter *Interpreter) getVariableValue(name string) string {
	value, exists := interpreter.VariableMap[name]
	if !exists {
		panic("Undefined variable: " + name)
	}

	switch v := value.(type) {
	case []lexer.Token: //Handle slice of tokens
		for _, token := range v {
			switch token.Tokentype {
			case lexer.IDENTIFIER:
				return interpreter.getVariableValue(*token.Value) // Resolve identifier.
			case lexer.STRING_LIT:
				return *token.Value
			case lexer.INT_LIT:
				return *token.Value

			default:
				panic("cant acess this variable")
			}
		}
	}
	return "-1" // If nothing matches, return nil.

}

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
		switch n.Token.Tokentype {
		case lexer.IDENTIFIER:
			fmt.Println(interpreter.getVariableValue(*n.Token.Value))
		default:
			fmt.Println(strings.ReplaceAll(*n.Token.Value, `"`, ""))
		}

	case *parser.Block:
		for _, block := range n.Body {
			//add scoping logic here
			interpreter.VisitNode(block)
			//remove the scope from the stack here

		}
	case *parser.IfStatement:
		var leftcondition interface{}
		var rightcondition interface{}

		switch value := n.Condition.Left.(type) {
		case lexer.Token:
			if value.Tokentype == lexer.IDENTIFIER {
				leftcondition = interpreter.getVariableValue(*value.Value)
			} else {
				leftcondition = *value.Value
			}
		}
		switch value := n.Condition.Right.(type) {
		case lexer.Token:
			fmt.Println(*value.Value)
			if value.Tokentype == lexer.IDENTIFIER {
				rightcondition = interpreter.getVariableValue(*value.Value)
			} else {
				rightcondition = *value.Value
			}

		}
		if leftcondition == rightcondition {
			interpreter.VisitNode(n.Body)
		}
	case *parser.LoopStatement:
		var leftcondition interface{}
		var rightcondition interface{}
		switch value := n.Condition.Left.(type) {
		case lexer.Token:
			if value.Tokentype == lexer.IDENTIFIER {
				leftcondition = interpreter.getVariableValue(*value.Value)
			} else {
				leftcondition = *value.Value
			}
		}
		switch value := n.Condition.Right.(type) {
		case lexer.Token:
			fmt.Println(*value.Value)
			if value.Tokentype == lexer.IDENTIFIER {
				rightcondition = interpreter.getVariableValue(*value.Value)
			} else {
				rightcondition = *value.Value
			}

		}

		if leftcondition == rightcondition {
			interpreter.VisitNode(n.Body)
		}

	default:
		fmt.Printf("Unknown node type: %s\n", reflect.TypeOf(n).String())
	}
}
