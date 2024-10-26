package interpreter

import (
	"compileringo/internal/lexer"
	"compileringo/internal/parser"
	"fmt"
	"reflect"
	"strings"

	"github.com/expr-lang/expr"
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

func evaluateExpression(str string) string {
	expression := str
	result, err := expr.Eval(expression, nil)
	if err != nil {
		fmt.Println("string is ")
		fmt.Println(str)
	}

	return fmt.Sprintf("%v", result)
}

func (interpreter *Interpreter) getVariableValue(name string) string {
	value, exists := interpreter.VariableMap[name]
	var sb strings.Builder
	if !exists {
		panic("Undefined variable: " + name)
	}

	switch v := value.(type) {
	case []lexer.Token:
		for _, token := range v {
			switch token.Tokentype {
			case lexer.IDENTIFIER:
				resolvedValue := interpreter.getVariableValue(*token.Value)
				sb.WriteString(resolvedValue)
			case lexer.STRING_LIT:
				sb.WriteString(*token.Value)
			case lexer.INT_LIT:
				sb.WriteString(*token.Value)
			case lexer.PLUS:
				sb.WriteString("+")
			case lexer.MINUS:
				sb.WriteString("-")
			case lexer.DIVIDE:
				sb.WriteString("/")
			case lexer.MULTIPLY:
				sb.WriteString("*")
			default:
				panic("Can't access this variable")
			}
		}
	}
	return sb.String()
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
		value := evaluateExpression(interpreter.getVariableValue(n.Name))
		interpreter.VariableMap[n.Name] = []lexer.Token{{
			Tokentype: lexer.INT_LIT,
			Value:     &value,
		}}

	case *parser.VariableReasign:
		if _, exists := interpreter.VariableMap[n.Name]; exists {
			// Get the current value of the variable

			// Assuming n.Value is a slice of lexer.Token
			var newExpression strings.Builder
			// Loop over the tokens in n.Value
			for _, token := range n.Value {
				switch token.Tokentype {
				case lexer.PLUS:
					newExpression.WriteString(" + ")
				case lexer.MINUS:
					newExpression.WriteString(" - ")
				case lexer.MULTIPLY:
					newExpression.WriteString(" * ")
				case lexer.DIVIDE:
					newExpression.WriteString(" / ")
				case lexer.IDENTIFIER:
					// Resolve the identifier value
					newValue := interpreter.getVariableValue(*token.Value)
					newExpression.WriteString(newValue)
				case lexer.INT_LIT:
					newExpression.WriteString(*token.Value) // Assuming token.Value is a pointer to a string
				case lexer.STRING_LIT:
					newExpression.WriteString(*token.Value) // Assuming token.Value is a pointer to a string
				default:
					panic("Unsupported token type in variable reassignment")
				}
			}

			// Evaluate the new expression
			evaluatedValue := evaluateExpression(newExpression.String())

			// Update the VariableMap with the new evaluated value
			interpreter.VariableMap[n.Name] = []lexer.Token{{
				Tokentype: lexer.INT_LIT,
				Value:     &evaluatedValue,
			}}
			fmt.Println("Updated value of", n.Name, ":", evaluatedValue)
		} else {
			panic("Variable reassigned without initializing")
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
			interpreter.VisitNode(block)

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
		for{
			interpreter.VisitNode(n.Body)
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
				if value.Tokentype == lexer.IDENTIFIER {
					rightcondition = interpreter.getVariableValue(*value.Value)
				} else {
					rightcondition = *value.Value
				}
			}
      if leftcondition == rightcondition{
      break
      }
		}

	default:
		fmt.Printf("Unknown node type: %s\n", reflect.TypeOf(n).String())
	}
}
