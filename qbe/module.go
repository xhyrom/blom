package qbe

// A complete QBE module
type Module struct {
	Functions []Function
	Types     []TypeDefinition
	Data      []Data
}

func (m Module) String() string {
	result := ""

	for _, function := range m.Functions {
		result += function.String()
	}

	for _, typ := range m.Types {
		result += typ.String()
	}

	for _, data := range m.Data {
		result += data.String()
	}

	return result
}
