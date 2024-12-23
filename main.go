package main

import (
	"blom/interpreter"
	"blom/lexer"
	"blom/parser"
	"blom/tokens"
	"blom/transpiler/javascript"
	"blom/transpiler/lua"
	"fmt"
	"os"

	"github.com/gookit/goutil/dump"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file>\n", os.Args[0])
		os.Exit(1)
	}

	content, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %s\n", err)
		os.Exit(1)
	}

	lexer := lexer.New(args[0], string(content))
	tkns := make([]tokens.Token, 0)

	c := lexer.Next()
	for c.Kind != tokens.Eof {
		tkns = append(tkns, *c)
		c = lexer.Next()
	}

	dump.Println(tkns)

	pp := parser.New(args[0])
	ast := pp.AST(args[0], string(content))

	dump.Config(func(o *dump.Options) {
		o.MaxDepth = 10
	})

	dump.Println(ast)

	inp := interpreter.New()
	dump.Println(inp.Interpret(ast))

	pp = parser.New(args[0])
	ast = pp.AST(args[0], string(content))
	transp := lua.LuaTranspiler{}
	code, _ := transp.Transpile(ast)

	fmt.Printf("%s\n", code)

	pp = parser.New(args[0])
	ast = pp.AST(args[0], string(content))
	transpp := javascript.JavascriptTranspiler{}
	code, _ = transpp.Transpile(ast)

	fmt.Printf("%s\n", code)
}
