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
	VariableMap []map[string]interface{}
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		FunctionMap: make(map[string]*parser.FunctionDeclaration),
		VariableMap: []map[string]interface{}{},
	}
}

func evaluateExpression(str string) (string, error) {
	expression := str
	result, err := expr.Eval(expression, nil)
	if err != nil {
		return str, err
	}

	return fmt.Sprintf("%v", result), nil
}
func (interpreter *Interpreter) PushStack() {
	interpreter.VariableMap = append(interpreter.VariableMap, make(map[string]interface{}))
}
func (interpreter *Interpreter) popStack() {
	if len(interpreter.VariableMap)-1 > 0 {
		interpreter.VariableMap = interpreter.VariableMap[:len(interpreter.VariableMap)-1]
	}
}

func (interpreter *Interpreter) getVariableValue(name string) string {
	var sb strings.Builder

	for i := len(interpreter.VariableMap) - 1; i >= 0; i-- {
		value, exists := interpreter.VariableMap[i][name]
		if !exists {
			continue
		}

		switch v := value.(type) {
		case []lexer.Token:
			for _, token := range v {
				switch token.Tokentype {
				case lexer.IDENTIFIER:
					resolvedValue := interpreter.getVariableValue(*token.Value)
					sb.WriteString(resolvedValue)
				case lexer.STRING_LIT, lexer.INT_LIT:
					sb.WriteString(*token.Value)
				case lexer.PLUS, lexer.MINUS, lexer.DIVIDE, lexer.MULTIPLY:
					sb.WriteString(*token.Value)
				default:
					panic("Can't access this variable")
				}
			}
			break
		}
	}

	if sb.Len() == 0 {
		panic(fmt.Sprintf("Variable '%s' is not declared but is being accessed", name))
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
		if _, exists := interpreter.VariableMap[len(interpreter.VariableMap)-1][n.Name]; exists {
			panic("cannot declare a variable twice")
		}
		for i := len(interpreter.VariableMap) - 1; i >= 0; i-- {
			interpreter.VariableMap[i][n.Name] = n.Value

			value, err := evaluateExpression(interpreter.getVariableValue(n.Name))
			if err != nil {
				interpreter.VariableMap[i][n.Name] = []lexer.Token{{
					Tokentype: lexer.INT_LIT,
					Value:     &value,
				}}
				_, exists := interpreter.VariableMap[i][n.Name]
				if exists {
					break
				}
			} else {
				interpreter.VariableMap[i][n.Name] = []lexer.Token{{
					Tokentype: lexer.INT_LIT,
					Value:     &value,
				}}
			}
			_, exists := interpreter.VariableMap[i][n.Name]
			if exists {
				break
			}
		}
	case *parser.VariableReasign:
		var newExpression strings.Builder
		for i := len(interpreter.VariableMap) - 1; i >= 0; i-- {
			if _, exists := interpreter.VariableMap[i][n.Name]; exists {
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
						newValue := interpreter.getVariableValue(*token.Value)
						newExpression.WriteString(newValue)
					case lexer.INT_LIT:
						newExpression.WriteString(*token.Value)
					case lexer.STRING_LIT:
						newExpression.WriteString(*token.Value)
					default:
						panic("Unsupported token type in variable reassignment")
					}
				}
				evaluatedValue, err := evaluateExpression(newExpression.String())

				if err != nil {
					interpreter.VariableMap[i][n.Name] = []lexer.Token{{
						Tokentype: lexer.INT_LIT,
						Value:     &evaluatedValue,
					}}
					_, exists := interpreter.VariableMap[i][n.Name]
					if exists {
						break
					}
				} else {
					interpreter.VariableMap[i][n.Name] = []lexer.Token{{
						Tokentype: lexer.INT_LIT,
						Value:     &evaluatedValue,
					}}
					_, exists := interpreter.VariableMap[i][n.Name]
					if exists {
						break
					}
				}

			}
		}
		if newExpression.Len() == 0 {
			panic("batman")
		}

	case *parser.PrintStatement:
		switch n.Token.Tokentype {
		case lexer.IDENTIFIER:
			fmt.Println(interpreter.getVariableValue(*n.Token.Value))
		default:
			fmt.Println(strings.ReplaceAll(*n.Token.Value, `"`, ""))
		}

	case *parser.Block:
		interpreter.PushStack()
		for _, block := range n.Body {
			interpreter.VisitNode(block)
		}
		interpreter.popStack()
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
		for {
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
			if leftcondition == rightcondition {
				break
			}
		}

	default:
		fmt.Printf("Unknown node type: %s\n", reflect.TypeOf(n).String())
	}
}
