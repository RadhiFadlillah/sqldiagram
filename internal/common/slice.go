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

func DeleteSliceItem[T any](items []T, index int) []T {
	var zero T
	if index < 0 || index >= len(items) {
		return items
	}

	copy(items[index:], items[index+1:])
	items[len(items)-1] = zero
	items = items[:len(items)-1]
	return items
}

func InsertSliceItem[T any](items []T, index int, newItem T) []T {
	var zero T
	if index < 0 || index >= len(items) {
		return items
	}

	items = append(items, zero)
	copy(items[index+1:], items[index:])
	items[index] = newItem
	return items
}
