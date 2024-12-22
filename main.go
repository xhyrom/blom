package main

import (
	"blom/lexer"
	"blom/parser"
	"blom/tokens"
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
	c := lexer.Next()
	for c.Kind != tokens.Eof {
		fmt.Println(c.Kind, c.Location, c.Value)

		c = lexer.Next()
	}

	parser := parser.New(args[0])
	ast := parser.AST(args[0], string(content))

	dump.Config(func(o *dump.Options) {
		o.MaxDepth = 10
	})

	dump.Println(ast)
}
