package consumers

func ConsumeComment(lex Lexer) {
	char := lex.CurrentChar()
	for char != '\n' {
		err := lex.Advance()

		if err != nil {
			break
		}

		char = lex.CurrentChar()
	}

	lex.NewLine()
}
