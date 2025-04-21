package main

import (
	"blom/lexer"
	"blom/tokens"
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/gookit/goutil/dump"
)

func main() {
	dump.Config(func(o *dump.Options) {
		o.MaxDepth = 100
	})

	compact := flag.Bool("compact", false, "Display tokens in a single-line format")
	pretty := flag.Bool("pretty", false, "Display tokens in a pretty format with one token per line")
	noColor := flag.Bool("no-color", false, "Disable colored output")
	help := flag.Bool("help", false, "Display help information")
	version := flag.Bool("version", false, "Display version information")
	readStdin := flag.Bool("stdin", false, "Read from standard input instead of files")
	flag.Parse()

	if *noColor {
		color.NoColor = true
	}

	if *help {
		printHelp()
		return
	}

	if *version {
		fmt.Println("blomray v1.0 - The Blom Token Illuminator")
		fmt.Println("Part of the Blom Programming Language Toolchain")
		return
	}

	var format string
	if *pretty {
		format = "pretty"
	} else if *compact {
		format = "compact"
	} else {
		format = "default"
	}

	args := flag.Args()

	if len(args) == 0 && !*readStdin {
		fmt.Fprintln(os.Stderr, "Error: No input file specified")
		fmt.Fprintln(os.Stderr, "Run 'blomray --help' for usage information")
		os.Exit(1)
	}

	if *readStdin {
		processStdin(format)
	} else {
		for _, inputFile := range args {
			if !strings.HasSuffix(inputFile, ".blom") {
				fmt.Fprintf(os.Stderr, "Warning: '%s' does not have .blom extension\n", inputFile)
			}
			processFile(inputFile, format)
		}
	}
}

func processStdin(format string) {
	scanner := bufio.NewScanner(os.Stdin)
	content := strings.Builder{}

	for scanner.Scan() {
		content.WriteString(scanner.Text() + "\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
		os.Exit(1)
	}

	lex := lexer.New("<stdin>", content.String())
	processTokens(lex, format)
}

func processFile(inputFile string, format string) {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	lex := lexer.New(inputFile, string(content))
	processTokens(lex, format)
}

func processTokens(lex *lexer.Lexer, format string) {
	tkns := make([]tokens.Token, 0)

	current := lex.Next()
	for current.Kind != tokens.Eof {
		tkns = append(tkns, *current)
		current = lex.Next()
	}

	switch format {
	case "compact":
		for _, token := range tkns {
			fmt.Printf("%s(row=%d, col=%d",
				token.Kind,
				token.Location.Row,
				token.Location.Column)

			if token.Value != "" {
				fmt.Printf(", value=%s", token.Value)
			}

			fmt.Print(") ")
		}
		fmt.Println()

	case "pretty":
		printPrettyTokens(tkns)

	default:
		dump.Println(tkns)
	}
}

func printPrettyTokens(tkns []tokens.Token) {
	indexStyle := color.New(color.FgHiBlack)
	kindStyle := color.New(color.FgHiCyan)
	locationStyle := color.New(color.FgHiBlack)

	valueStyles := map[tokens.TokenKind]*color.Color{
		tokens.Identifier:     color.New(color.FgHiWhite, color.Bold),
		tokens.StringLiteral:  color.New(color.FgHiGreen),
		tokens.IntLiteral:     color.New(color.FgHiMagenta),
		tokens.FloatLiteral:   color.New(color.FgHiMagenta),
		tokens.BooleanLiteral: color.New(color.FgHiRed),
	}

	maxIndexWidth := 0
	maxKindWidth := 0
	maxLocationWidth := 0
	hasValues := false

	for i, token := range tkns {
		indexWidth := len(fmt.Sprintf("[%d]", i))
		if indexWidth > maxIndexWidth {
			maxIndexWidth = indexWidth
		}

		kindWidth := len(token.Kind.String())
		if kindWidth > maxKindWidth {
			maxKindWidth = kindWidth
		}

		locationWidth := len(fmt.Sprintf("row=%d col=%d", token.Location.Row, token.Location.Column))
		if locationWidth > maxLocationWidth {
			maxLocationWidth = locationWidth
		}

		if token.Value != "" {
			hasValues = true
		}
	}

	for i, token := range tkns {
		indexStr := indexStyle.Sprintf("[%d]", i)
		indexPadding := strings.Repeat(" ", maxIndexWidth-len(fmt.Sprintf("[%d]", i)))

		var kindStr string
		switch token.Kind {
		case tokens.Fun, tokens.Return, tokens.Val, tokens.Var, tokens.If, tokens.Else, tokens.For, tokens.While:
			kindStr = color.New(color.FgHiBlue).Sprint(token.Kind.String())
		case tokens.Assign, tokens.Equals, tokens.NotEquals, tokens.Plus, tokens.Minus, tokens.Asterisk, tokens.Slash,
			tokens.LessThan, tokens.LessThanOrEqual, tokens.GreaterThan, tokens.GreaterThanOrEqual:
			kindStr = color.New(color.FgHiYellow).Sprint(token.Kind.String())
		default:
			kindStr = kindStyle.Sprint(token.Kind.String())
		}

		kindPadding := strings.Repeat(" ", maxKindWidth-len(token.Kind.String()))

		locStr := locationStyle.Sprintf("row=%d col=%d", token.Location.Row, token.Location.Column)
		locPadding := strings.Repeat(" ", maxLocationWidth-len(fmt.Sprintf("row=%d col=%d", token.Location.Row, token.Location.Column)))

		var valStr string
		if token.Value != "" {
			style, ok := valueStyles[token.Kind]
			if !ok {
				style = color.New(color.FgWhite)
			}
			valStr = fmt.Sprintf("= %s", style.Sprint(token.Value))
		} else if hasValues {
			valStr = ""
		}

		fmt.Printf("%s%s %s%s @ %s%s %s\n",
			indexPadding, indexStr,
			kindStr, kindPadding,
			locStr, locPadding,
			valStr)
	}
}

func printHelp() {
	fmt.Println("blomray - Blom Token Illuminator")
	fmt.Println()
	fmt.Println("Usage: blomray [options] [file.blom...]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -compact     Display tokens in a compact single-line format")
	fmt.Println("  -pretty      Display tokens in a pretty format with one token per line")
	fmt.Println("  -no-color    Disable colored output (for pretty format)")
	fmt.Println("  -stdin       Read from standard input instead of files")
	fmt.Println("  -version     Display version information")
	fmt.Println("  -help        Display this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  blomray sample.blom                        # Display tokens in default format")
	fmt.Println("  blomray -pretty sample.blom                # Display tokens in pretty, colorful format")
	fmt.Println("  blomray -compact sample.blom               # Display tokens in compact format")
	fmt.Println("  echo 'val x = 5' | blomray -stdin -pretty  # Process code from stdin")
	fmt.Println()
	fmt.Println("Like a ray of light illuminates particles in the air,")
	fmt.Println("blomray reveals the tokens in your Blom source code.")
}
