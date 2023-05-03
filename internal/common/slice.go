package common

func FindSlice[T any](items []T, finder func(item T) bool) (T, int) {
	for i, item := range items {
		if finder(item) {
			return item, i
		}
	}

	var zero T
	return zero, -1
}

func UniqueSlice[T comparable](items []T) []T {
	var uniqueList []T
	tracker := NewSet[T]()
	for _, item := range items {
		if !tracker.Has(item) {
			uniqueList = append(uniqueList, item)
			tracker.Put(item)
		}
	}
	return uniqueList
}
