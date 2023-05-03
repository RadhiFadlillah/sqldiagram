package common

import "sort"

type OrderedMap[K comparable, V any] struct {
	entries  map[K]V
	orders   map[K]int
	maxOrder int
}

func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		entries: make(map[K]V),
		orders:  make(map[K]int),
	}
}

func (m *OrderedMap[K, V]) init() {
	if m == nil {
		m = NewOrderedMap[K, V]()
	}
	if m.entries == nil {
		m.entries = make(map[K]V)
	}
	if m.orders == nil {
		m.orders = make(map[K]int)
	}
}

func (m OrderedMap[K, V]) Has(k K) bool {
	_, entryExist := m.entries[k]
	_, orderExist := m.orders[k]
	return entryExist && orderExist
}

// Put insert entry to map, but will not change the order if it already exist before.
func (m *OrderedMap[K, V]) Put(k K, v V) {
	m.init()
	m.entries[k] = v
	if !m.Has(k) {
		m.orders[k] = m.maxOrder
		m.maxOrder++
	}
}

// Replace insert entry to map, and will change the order.
func (m *OrderedMap[K, V]) Replace(k K, v V) {
	m.init()
	m.entries[k] = v
	m.orders[k] = m.maxOrder
	m.maxOrder++
}

func (m OrderedMap[K, V]) Get(k K) V {
	return m.entries[k]
}

func (m OrderedMap[K, V]) Keys() []K {
	var keys []K
	for k := range m.entries {
		keys = append(keys, k)
	}
	return keys
}

func (m OrderedMap[K, V]) Values() []V {
	var kvs []struct {
		Key   K
		Value V
	}
	for k, v := range m.entries {
		kvs = append(kvs, struct {
			Key   K
			Value V
		}{k, v})
	}

	sort.Slice(kvs, func(i int, j int) bool {
		return m.orders[kvs[i].Key] < m.orders[kvs[j].Key]
	})

	values := make([]V, len(kvs))
	for i, kv := range kvs {
		values[i] = kv.Value
	}

	return values
}

func (m OrderedMap[K, V]) ForEach(cb func(K, V)) {
	for k, v := range m.entries {
		cb(k, v)
	}
}

func (m OrderedMap[K, V]) Size() int {
	return len(m.entries)
}
