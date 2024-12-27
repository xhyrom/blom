package debug

import (
	"blom/ast"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var fileCache map[string][]byte

type SourceLocation struct {
	File   string
	Row    uint64
	Column uint64
}

func (s *SourceLocation) Copy() SourceLocation {
	return SourceLocation{
		File:   s.File,
		Row:    s.Row,
		Column: s.Column,
	}
}

type Hint struct {
	Msg  string
	Code string
}

func NewHint(msg string, code string) *Hint {
	return &Hint{
		Msg:  msg,
		Code: code,
	}
}

func NewSourceLocation(file string, row uint64, column uint64) SourceLocation {
	if fileCache == nil {
		fileCache = map[string][]byte{}
	}
	return SourceLocation{
		File:   file,
		Row:    row,
		Column: column,
	}
}

func NewSourceLocationFromExpression(file string, expression ast.Expression) SourceLocation {
	if fileCache == nil {
		fileCache = map[string][]byte{}
	}

	row := expression.Location().Row
	column := expression.Location().Column

	return SourceLocation{
		File:   file,
		Row:    row,
		Column: column + 1,
	}
}

func (s *SourceLocation) Stringify() string {
	return fmt.Sprintf("#line %d \"%s\"", s.Row, s.File)
}

func (s *SourceLocation) ThrowError(msg string, top bool, hints ...*Hint) {
	if !top {
		fmt.Fprintln(os.Stderr, "──────────────────────────────────────")
	}

	n, _ := fmt.Fprintf(os.Stderr, "[%s:%d:%d]", filepath.Base(s.File), s.Row, s.Column)
	fmt.Fprintf(os.Stderr, " \x1b[31;1mError\x1b[0m: %s\n", msg)

	code, ok := fileCache[s.File]

	if !ok {
		code, _ = os.ReadFile(s.File)
		fileCache[s.File] = code
	}

	lines := strings.Split(string(code), "\n")

	fmt.Fprintln(os.Stderr, "━━━━━━━━━━━━━━━━ code ━━━━━━━━━━━━━━━━")

	startLine := int64(s.Row - 5)
	if startLine < 0 {
		startLine = 0
	}

	width := len(fmt.Sprintf("%d", s.Row))

	if len(hints) == 0 {
		for i := uint64(startLine); i < s.Row; i++ {
			fmt.Fprintf(os.Stderr, "%*d | %s\n", width, i+1, lines[i])
		}

		fmt.Fprintf(os.Stderr, "%*c |%*c", width, ' ', s.Column, ' ')
		fmt.Fprintf(os.Stderr, "^\n")
	} else {
		for _, hint := range hints {
			if hint != nil && len(hint.Code) > 0 {
				for i := uint64(startLine); i < s.Row-1; i++ {
					fmt.Fprintf(os.Stderr, "%*d | %s\n", width, i+1, lines[i])
				}

				fmt.Fprintf(os.Stderr, "%*d | %s\x1b[32;1m%s\x1b[0m%s\n", width, s.Row, lines[s.Row-1][:s.Column-1], hint.Code, lines[s.Row-1][s.Column-1:])

				fmt.Fprintf(os.Stderr, "%*c |%*c", width, ' ', s.Column, ' ')
				fmt.Fprintf(os.Stderr, "\x1b[32;1m%s\x1b[0m\n", strings.Repeat("+", len(hint.Code)))

				fmt.Fprintf(os.Stderr, "%*c\x1b[34;1mHint\x1b[0m: %s\n", n, ' ', hint.Msg)

				if hint != hints[len(hints)-1] {
					fmt.Fprintln(os.Stderr, "")
				}
			} else {
				for i := uint64(startLine); i < s.Row; i++ {
					fmt.Fprintf(os.Stderr, "%*d | %s\n", width, i+1, lines[i])
				}
				fmt.Fprintf(os.Stderr, "%*c |%*c", width, ' ', s.Column, ' ')
				fmt.Fprintf(os.Stderr, "^\n")
			}
		}
	}

	fmt.Fprintln(os.Stderr, "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	os.Exit(1)
}

func (s *SourceLocation) ThrowWarning(msg string, top bool, hint *Hint) {
	if !top {
		fmt.Fprintln(os.Stderr, "──────────────────────────────────────")
	}

	if hint != nil {
		n, _ := fmt.Fprintf(os.Stderr, "[%s:%d:%d]", filepath.Base(s.File), s.Row, s.Column)
		fmt.Fprintf(os.Stderr, " \x1b[33;1mWarning\x1b[0m: %s\n", msg)
		fmt.Fprintf(os.Stderr, "%*c \x1b[34;1mHint\x1b[0m: %s\n", n, ' ', hint.Msg)
	} else {
		fmt.Fprintf(os.Stderr, "[%s:%d:%d] \x1b[33;1Warning\x1b[0m: %s\n", filepath.Base(s.File), s.Row, s.Column, msg)
	}

	code, ok := fileCache[s.File]

	if !ok {
		code, _ = os.ReadFile(s.File)

		fileCache[s.File] = code
	}

	lines := strings.Split(string(code), "\n")

	fmt.Fprintln(os.Stderr, "━━━━━━━━━━━━━━━━ code ━━━━━━━━━━━━━━━━")

	startLine := int64(s.Row - 5)
	if startLine < 0 {
		startLine = 0
	}

	width := len(fmt.Sprintf("%d", s.Row))

	if hint != nil && len(hint.Code) > 0 {
		for i := uint64(startLine); i < s.Row-1; i++ {
			fmt.Fprintf(os.Stderr, "%*d | %s\n", width, i+1, lines[i])
		}
		fmt.Fprintf(os.Stderr, "%*d | %s\x1b[32;1m%s\x1b[0m%s\n", width, s.Row, lines[s.Row-1][:s.Column-1], hint.Code, lines[s.Row-1][s.Column-1:])

		fmt.Fprintf(os.Stderr, "%*c |%*c", width, ' ', s.Column, ' ')
		fmt.Fprintf(os.Stderr, "\x1b[32;1m%s\x1b[0m\n", strings.Repeat("+", len(hint.Code)))
	} else {
		for i := uint64(startLine); i < s.Row; i++ {
			fmt.Fprintf(os.Stderr, "%*d | %s\n", width, i+1, lines[i])
		}
		fmt.Fprintf(os.Stderr, "%*c |%*c", width, ' ', s.Column, ' ')
		fmt.Fprintf(os.Stderr, "^\n")
	}
	fmt.Fprintln(os.Stderr, "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
}

func LogError(msg string, hint *Hint) {
	if hint != nil {
		fmt.Fprintf(os.Stderr, "\x1b[31;1mError\x1b[0m: %s\n", msg)
		fmt.Fprintf(os.Stderr, "\x1b[34;1mHint\x1b[0m: %s\n", hint.Msg)
	} else {
		fmt.Fprintf(os.Stderr, "\x1b[31;1mError\x1b[0m: %s\n", msg)
	}
}

func LogWarning(msg string, hint *Hint) {
	if hint != nil {
		fmt.Fprintf(os.Stderr, "\x1b[33;1mWarning\x1b[0m: %s\n", msg)
		fmt.Fprintf(os.Stderr, "\x1b[34;1mHint\x1b[0m: %s\n", hint.Msg)
	} else {
		fmt.Fprintf(os.Stderr, "\x1b[33;1Warning\x1b[0m: %s\n", msg)
	}
}
