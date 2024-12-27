package cli

import (
	"blom/interpreter"
	"blom/lexer"
	"blom/parser"
	"blom/tokens"
	"os"
	"strings"

	"github.com/gookit/goutil/dump"
)

func Run(args []string) {
	var emitTokens bool
	var emitAst bool
	var inputFile string

	for _, arg := range args {
		switch arg {
		case "-tokens", "--emit-tokens":
			emitTokens = true
		case "-ast", "--emit-ast":
			emitAst = true
		default:
			if strings.HasSuffix(arg, ".blom") {
				inputFile = arg
			}
		}
	}

	if inputFile == "" {
		panic("No input file specified")
	}

	content, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	lex := lexer.New(inputFile, string(content))
	tkns := make([]tokens.Token, 0)

	current := lex.Next()
	for current.Kind != tokens.Eof {
		tkns = append(tkns, *current)
		current = lex.Next()
	}

	if emitTokens {
		dump.Println(tkns)
	}

	parser := parser.New(inputFile)
	ast := parser.AST(inputFile, string(content))

	if emitAst {
		dump.Println(ast)
	}

	inp := interpreter.New(inputFile)
	dump.Println(inp.Interpret(ast))
}
