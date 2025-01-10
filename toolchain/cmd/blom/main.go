package main

import (
	"blom/analyzer"
	"blom/interpreter"
	"blom/lexer"
	"blom/parser"
	"blom/tokens"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gookit/goutil/dump"
)

func main() {
	dump.Config(func(o *dump.Options) {
		o.MaxDepth = 100
	})

	args := os.Args[1:]

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

	analyzer := analyzer.New(inputFile, ast)
	analyzer.Analyze()

	if emitAst {
		dump.Println(ast)
	}

	inp := interpreter.New()

	fmt.Printf("Interpreting %s\n", inputFile)

	startTime := time.Now()
	inp.Interpret(ast)
	endTime := time.Since(startTime)

	fmt.Printf("Interpreted %s ran for %s\n", inputFile, endTime)
}
