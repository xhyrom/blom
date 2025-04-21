package main

import (
	"blom/mangling"
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	noParams := flag.Bool("no-params", false, "Don't display function parameters")
	noTypes := flag.Bool("no-types", false, "Don't display parameter types")
	compact := flag.Bool("compact", false, "Display in compact format")
	showErrors := flag.Bool("show-errors", false, "Show errors during demangling")
	help := flag.Bool("help", false, "Display help information")
	version := flag.Bool("version", false, "Display version information")
	flag.Parse()

	if *help {
		printHelp()
		return
	}

	if *version {
		fmt.Println("blomprism v1.0 - The Blom Symbol Demangler")
		fmt.Println("Part of the Blom Programming Language Toolchain")
		return
	}

	decoder := mangling.NewDecoder()
	args := flag.Args()

	if len(args) > 0 {
		for _, arg := range args {
			processName(arg, decoder, *noParams, *noTypes, *compact, *showErrors)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line != "" {
				processName(line, decoder, *noParams, *noTypes, *compact, *showErrors)
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "error: reading from stdin: %v\n", err)
			os.Exit(1)
		}
	}
}

func processName(name string, decoder *mangling.Decoder, noParams bool, noTypes bool, compact bool, showErrors bool) {
	if !strings.HasPrefix(name, "_Z") {
		fmt.Println(name)
		return
	}

	demangled, err := decoder.DemangleSymbol(name)
	if err != nil {
		if showErrors {
			fmt.Fprintf(os.Stderr, "error: demangling symbol '%s': %v\n", name, err)
		}

		fmt.Println(name)
		return
	}

	var result string

	if compact {
		result = demangled.Name
	} else {
		result = demangled.Name

		if !noParams {
			result += "("
			if !noTypes && len(demangled.Parameters) > 0 {
				result += strings.Join(demangled.Parameters, ", ")
			}
			result += ")"
		}

		if !noTypes {
			result += " -> " + demangled.ReturnType
		}
	}

	fmt.Println(result)
}

func printHelp() {
	fmt.Println("blomprism - Blom Symbol Demangler")
	fmt.Println()
	fmt.Println("Usage: blomprism [options] [symbols...]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -no-params    Don't display function parameters")
	fmt.Println("  -no-types     Don't display parameter types")
	fmt.Println("  -compact      Display only function names (no parameters or return types)")
	fmt.Println("  -show-errors  Show errors during demangling")
	fmt.Println("  -version      Display version information")
	fmt.Println("  -help         Display this help message")
	fmt.Println()
	fmt.Println("If no symbols are specified, blomprism reads from standard input.")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  blomprism _Z3addii_Ri                # Demangle a single symbol")
	fmt.Println("  echo \"_Z3addii_Ri\" | blomprism       # Pipe input from another command")
	fmt.Println("  blomprism -compact _Z3addii_Ri       # Show just the function name")
	fmt.Println()
	fmt.Println("Like a prism reveals the spectrum of colors in light,")
	fmt.Println("blomprism reveals the original function signature in mangled symbols.")
}
