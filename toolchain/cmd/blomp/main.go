package main

import (
	"blom/lexer"
	"blom/parser"
	"blom/tokens"
	"fmt"
	"os"
	"strings"

	"github.com/gookit/goutil/dump"
)

func main() {
	dump.Config(func(o *dump.Options) {
		o.MaxDepth = 100
	})

	args := os.Args[1:]

	var emitTokens bool

	var inputFile string

	for _, arg := range args {
		switch arg {
		case "-tokens", "--emit-tokens":
			emitTokens = true
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

	fmt.Println("Abstract Syntax Tree before analysis:")
	dump.Println(ast)

	/*analyzer := analyzer.New(inputFile, ast)
	analyzer.Analyze()

	fmt.Println("Abstract Syntax Tree after analysis:")
	dump.Println(ast)*/
}
