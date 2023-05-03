package common

type Map[K comparable, V any] map[K]V

func NewMap[K comparable, V any]() Map[K, V] {
	return make(Map[K, V])
}

func (m Map[K, V]) Put(k K, v V) {
	m[k] = v
}

func (m Map[K, V]) Get(k K) V {
	return m[k]
}

func (m Map[K, V]) Has(k K) bool {
	_, exist := m[k]
	return exist
}

func (m Map[K, V]) Keys() []K {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func (m Map[K, V]) Values() []V {
	var values []V
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
