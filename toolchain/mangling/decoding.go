package mangling

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type DecodedFunction struct {
	Name       string
	Parameters []string
	ReturnType string
}

type Decoder struct{}

func NewDecoder() *Decoder {
	return &Decoder{}
}

func (d *Decoder) DemangleSymbol(mangled string) (*DecodedFunction, error) {
	if !strings.HasPrefix(mangled, "_Z") {
		return nil, errors.New("not a mangled symbol: missing _Z prefix")
	}

	input := mangled[2:] // skip _Z prefix

	name, remaining, err := d.decodeName(input)
	if err != nil {
		return nil, err
	}

	returnSepIdx := strings.Index(remaining, "_R")
	if returnSepIdx == -1 {
		return nil, errors.New("malformed symbol: missing return type separator")
	}

	paramStr := remaining[:returnSepIdx]
	returnStr := remaining[returnSepIdx+2:] // skip _R prefix

	params, err := d.decodeParameterTypes(paramStr)
	if err != nil {
		return nil, err
	}

	returnType, _, err := d.decodeType(returnStr)
	if err != nil {
		return nil, err
	}

	return &DecodedFunction{
		Name:       name,
		Parameters: params,
		ReturnType: returnType,
	}, nil
}

func (d *Decoder) decodeName(input string) (string, string, error) {
	if len(input) == 0 {
		return "", "", errors.New("empty input")
	}

	if input[0] == 'N' {
		input = input[1:] // skip 'N'
		var components []string

		// parse each name component
		for len(input) > 0 && input[0] != 'E' {
			if !isDigit(input[0]) {
				return "", "", errors.New("expected digit in name length")
			}

			lengthEnd := 0
			for lengthEnd < len(input) && isDigit(input[lengthEnd]) {
				lengthEnd++
			}

			length, err := strconv.Atoi(input[:lengthEnd])
			if err != nil {
				return "", "", err
			}

			if lengthEnd+length > len(input) {
				return "", "", errors.New("name component exceeds input length")
			}

			components = append(components, input[lengthEnd:lengthEnd+length])
			input = input[lengthEnd+length:]
		}

		if len(input) == 0 || input[0] != 'E' {
			return "", "", errors.New("missing 'E' terminator for nested name")
		}

		return strings.Join(components, "::"), input[1:], nil // Skip 'E'
	}

	if !isDigit(input[0]) {
		return "", "", errors.New("expected digit in name length")
	}

	lengthEnd := 0
	for lengthEnd < len(input) && isDigit(input[lengthEnd]) {
		lengthEnd++
	}

	length, err := strconv.Atoi(input[:lengthEnd])
	if err != nil {
		return "", "", err
	}

	if lengthEnd+length > len(input) {
		return "", "", errors.New("name exceeds input length")
	}

	name := input[lengthEnd : lengthEnd+length]
	return name, input[lengthEnd+length:], nil
}

func (d *Decoder) decodeParameterTypes(input string) ([]string, error) {
	if input == "v" {
		return []string{}, nil // void parameter list
	}

	var params []string
	remaining := input

	for len(remaining) > 0 {
		param, rest, err := d.decodeType(remaining)
		if err != nil {
			return nil, err
		}

		params = append(params, param)
		remaining = rest
	}

	return params, nil
}

func (d *Decoder) decodeType(input string) (string, string, error) {
	if len(input) == 0 {
		return "", "", errors.New("empty type")
	}

	firstChar := input[0]
	remaining := input[1:] // skip type code

	// handle basic types
	if typeStr, ok := ReverseTypeCodeMap[TypeCode(firstChar)]; ok {
		return typeStr, remaining, nil
	}

	// handle pointers
	if firstChar == 'P' {
		pointee, rest, err := d.decodeType(remaining)
		if err != nil {
			return "", "", err
		}
		return "*" + pointee, rest, nil
	}

	// handle function types
	if firstChar == 'F' {
		returnType, afterReturn, err := d.decodeType(remaining)
		if err != nil {
			return "", "", err
		}

		var params []string
		current := afterReturn

		for len(current) > 0 && current[0] != 'E' {
			param, rest, err := d.decodeType(current)
			if err != nil {
				return "", "", err
			}
			params = append(params, param)
			current = rest
		}

		if len(current) == 0 || current[0] != 'E' {
			return "", "", errors.New("missing 'E' terminator for function type")
		}

		paramsStr := strings.Join(params, ", ")
		return fmt.Sprintf("fn(%s) -> %s", paramsStr, returnType), current[1:], nil
	}

	// handle custom types
	if firstChar == 'X' {
		lengthEnd := 0
		for lengthEnd < len(remaining) && isDigit(remaining[lengthEnd]) {
			lengthEnd++
		}

		if lengthEnd == 0 {
			return "", "", errors.New("expected length for custom type")
		}

		length, err := strconv.Atoi(remaining[:lengthEnd])
		if err != nil {
			return "", "", err
		}

		if lengthEnd+length > len(remaining) {
			return "", "", errors.New("custom type name exceeds input")
		}

		typeName := remaining[lengthEnd : lengthEnd+length]
		return typeName, remaining[lengthEnd+length:], nil
	}

	return "", "", fmt.Errorf("unknown type code: %c", firstChar)
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
