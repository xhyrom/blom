package native

import (
	"blom/interpreter/objects"
	"fmt"
	"strings"
)

func Printf(parameters []objects.Object) interface{} {
	format := parameters[0].(*objects.StringObject).Value().(string)
	args := make([]interface{}, len(parameters)-1)

	for i, parameter := range parameters[1:] {
		switch parameter := parameter.(type) {
		case *objects.IntObject:
			args[i] = parameter.Value()
		case *objects.FloatObject:
			args[i] = parameter.Value()
		case *objects.DoubleObject:
			args[i] = parameter.Value()
		case *objects.StringObject:
			args[i] = unescape(parameter.Value().(string))
		case *objects.CharacterObject:
			args[i] = parameter.Value()
		case *objects.BooleanObject:
			val := parameter.Value().(bool)
			if val {
				args[i] = 1
			} else {
				args[i] = 0
			}
		default:
			panic(fmt.Sprintf("'%T' is not a valid parameter", parameter))
		}
	}

	fmt.Printf(unescape(format), args...)

	return nil
}

func unescape(text string) string {
	text = strings.ReplaceAll(text, "\\n", "\n")
	text = strings.ReplaceAll(text, "\\t", "\t")
	text = strings.ReplaceAll(text, "\\r", "\r")
	text = strings.ReplaceAll(text, "\\f", "\f")
	text = strings.ReplaceAll(text, "\\v", "\v")
	text = strings.ReplaceAll(text, "\\b", "\b")
	text = strings.ReplaceAll(text, "\\a", "\a")
	text = strings.ReplaceAll(text, "\\0", "\x00")
	text = strings.ReplaceAll(text, "\\\"", "\"")
	text = strings.ReplaceAll(text, "\\'", "'")
	text = strings.ReplaceAll(text, "\\\\", "\\")

	return text
}
