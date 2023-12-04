package set

type Set[T comparable] map[T]struct{}

func New[T comparable]() Set[T] {
	return make(map[T]struct{})
}

func (s Set[T]) Add(values ...T) Set[T] {
	for _, v := range values {
		s[v] = struct{}{}
	}
	return s
}

func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) ToSlice() []T {
	var result []T

	for value := range s {
		result = append(result, value)
	}

	return result
}
