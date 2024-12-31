package scope

type Scope[T any] map[string]T

func New[T any]() Scope[T] {
	return make(Scope[T])
}

func (s Scope[T]) Set(key string, value T) {
	s[key] = value
}

func (s Scope[T]) Get(key string) (T, bool) {
	val, ok := s[key]
	return val, ok
}

func (s Scope[T]) Exists(key string) bool {
	_, ok := s[key]
	return ok
}
