package main

import (
	"blom/lexer"
	"fmt"
	"os"
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
	for c != nil {
		fmt.Println(c.Kind, c.Location, c.Value)

		c = lexer.Next()
	}
}
