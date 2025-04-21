package main

import (
	"blom/analyzer"
	"blom/lexer"
	"blom/parser"
	"blom/tokens"
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/yassinebenaid/godump"
)

func main() {
	compact := flag.Bool("compact", false, "Display AST in a more compact format")
	analyze := flag.Bool("analyze", false, "Run semantic analysis before displaying AST")
	noColor := flag.Bool("no-color", false, "Disable colored output")
	help := flag.Bool("help", false, "Display help information")
	version := flag.Bool("version", false, "Display version information")
	readStdin := flag.Bool("stdin", false, "Read from standard input instead of files")
	flag.Parse()

	if *help {
		printHelp()
		return
	}

	if *version {
		fmt.Println("blomsight v1.0 - The Blom AST Visualizer")
		fmt.Println("Part of the Blom Programming Language Toolchain")
		return
	}

	var format string
	if *compact {
		format = "compact"
	} else {
		format = "default"
	}

	args := flag.Args()

	if len(args) == 0 && !*readStdin {
		fmt.Fprintln(os.Stderr, "Error: No input file specified")
		fmt.Fprintln(os.Stderr, "Run 'blomsight --help' for usage information")
		os.Exit(1)
	}

	if *readStdin {
		processStdin(format, *analyze, *noColor)
	} else {
		for _, inputFile := range args {
			if !strings.HasSuffix(inputFile, ".blom") {
				fmt.Fprintf(os.Stderr, "Warning: '%s' does not have .blom extension\n", inputFile)
			}

			processFile(inputFile, format, *analyze, *noColor)
		}
	}
}

func processStdin(format string, analyze bool, noColor bool) {
	scanner := bufio.NewScanner(os.Stdin)
	content := strings.Builder{}

	for scanner.Scan() {
		content.WriteString(scanner.Text() + "\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
		os.Exit(1)
	}

	processContent("<stdin>", content.String(), format, analyze, noColor)
}

func processFile(inputFile string, format string, analyze bool, noColor bool) {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	processContent(inputFile, string(content), format, analyze, noColor)
}

func processContent(filename string, content string, format string, analyze bool, noColor bool) {
	lex := lexer.New(filename, content)
	tkns := make([]tokens.Token, 0)

	current := lex.Next()
	for current.Kind != tokens.Eof {
		tkns = append(tkns, *current)
		current = lex.Next()
	}

	parser := parser.New(filename)
	ast := parser.AST(filename, content)

	if analyze {
		analyzer := analyzer.New(filename, ast)
		analyzer.Analyze()
	}

	var theme godump.Theme = godump.Theme{}
	if !noColor {
		theme = godump.DefaultTheme
	}

	switch format {
	case "compact":
		dumper := godump.Dumper{
			Theme: theme,
		}
		dumper.Println(ast)

	default:
		dumper := godump.Dumper{
			Theme:                   theme,
			ShowPrimitiveNamedTypes: true,
		}

		dumper.Println(ast)
	}
}

func printHelp() {
	fmt.Println("blomsight - Blom AST Visualizer")
	fmt.Println()
	fmt.Println("Usage: blomsight [options] [file.blom...]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -compact     Display AST in a more compact format")
	fmt.Println("  -pretty      Display AST in a tree-like structure")
	fmt.Println("  -analyze     Run semantic analysis before displaying AST")
	fmt.Println("  -no-color    Disable colored output")
	fmt.Println("  -stdin       Read from standard input instead of files")
	fmt.Println("  -version     Display version information")
	fmt.Println("  -help        Display this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  blomsight sample.blom                  # Display the AST in default format")
	fmt.Println("  blomsight -compact sample.blom         # Display the AST in a compact format")
	fmt.Println("  blomsight -analyze sample.blom         # Run analysis before displaying AST")
	fmt.Println("  echo 'fun main() {}' | blomsight -stdin  # Process code from stdin")
	fmt.Println()
	fmt.Println("Like sight reveals the hidden structures of the world,")
	fmt.Println("blomsight reveals the structure of your Blom program.")
}
