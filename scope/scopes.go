package scope

type Scopes[T any] []Scope[T]

func NewScopes[T any]() *Scopes[T] {
	s := make(Scopes[T], 0)
	return &s
}

func (s *Scopes[T]) Exists(key string) bool {
	for _, scope := range *s {
		if scope.Exists(key) {
			return true
		}
	}

	return false
}

func (s *Scopes[T]) GetValue(key string) (T, bool) {
	for i := len(*s) - 1; i >= 0; i-- {
		val, ok := (*s)[i].Get(key)
		if ok {
			return val, ok
		}
	}

	var zero T
	return zero, false
}

func (s *Scopes[T]) GetValueOrDefault(key string, defaultValue T) T {
	val, ok := s.GetValue(key)
	if ok {
		return val
	}

	return defaultValue
}

func (s *Scopes[T]) Set(key string, value T) {
	s.Last().Set(key, value)
}

func (s *Scopes[T]) Last() Scope[T] {
	return (*s)[len(*s)-1]
}

func (s *Scopes[T]) Append() {
	*s = append(*s, NewScope[T]())
}

func (s *Scopes[T]) Pop() Scope[T] {
	scope := s.Last()
	*s = (*s)[:len(*s)-1]
	return scope
}
