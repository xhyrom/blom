package main

import (
	"blom/analyzer"
	"blom/compiler"
	"blom/lexer"
	"blom/parser"
	"blom/tokens"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
	var emitSse bool

	var inputFile string
	var outputFile string = "a.out"

	for i, arg := range args {
		switch arg {
		case "-tokens", "--emit-tokens":
			emitTokens = true
		case "-ast", "--emit-ast":
			emitAst = true
		case "-o", "--output":
			outputFile = args[i+1]
			i += 1
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

	analyzer := analyzer.New(inputFile, ast)
	analyzer.Analyze()

	if emitAst {
		dump.Println(ast)
	}

	comp := compiler.New(compiler.QBE)
	sse := comp.Compile(ast)

	if emitSse {
		fmt.Println(sse)
	}

	outputDir := filepath.Dir(outputFile)
	if outputDir != "." {
		err := os.MkdirAll(outputDir, 0755)
		if err != nil {
			panic(err)
		}
	}

	sseFile := filepath.Join(outputDir, "out.sse")
	asmFile := filepath.Join(outputDir, "out.s")

	err = os.WriteFile(sseFile, []byte(sse), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Running %s\n", inputFile)
	startTime := time.Now()
	cmd := exec.Command("sh", "-c",
		fmt.Sprintf("qbe -o %s %s && cc -O3 %s -o %s",
			asmFile,    // QBE output assembly
			sseFile,    // QBE input
			asmFile,    // CC input
			outputFile, // Final executable
		))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	// Run the compiled executable
	cmd = exec.Command(outputFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
	endTime := time.Since(startTime)

	fmt.Printf("Compiled %s ran for %s\n", inputFile, endTime)

	os.Exit(cmd.ProcessState.ExitCode())
}
