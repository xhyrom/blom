package qbe

type ComparisonType int

const (
	LessThan ComparisonType = iota
	LessThanOrEqual
	GreaterThan
	GreaterThanOrEqual
	Equal
	NotEqual
)

func (c ComparisonType) String() string {
	switch c {
	case LessThan:
		return "lt"
	case LessThanOrEqual:
		return "le"
	case GreaterThan:
		return "gt"
	case GreaterThanOrEqual:
		return "ge"
	case Equal:
		return "eq"
	case NotEqual:
		return "ne"
	}

	panic("comparision type unreachable")
}
