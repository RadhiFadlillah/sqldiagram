package common

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) Put(keys ...T) {
	for _, k := range keys {
		s[k] = struct{}{}
	}
}

func (s Set[T]) Has(key T) bool {
	_, exist := s[key]
	return exist
}

func (s Set[T]) Keys() []T {
	var keys []T
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}
