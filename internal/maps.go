package internal

func GetKeys[T comparable, V any](m map[T]V) (keys []T) {
	for key := range m {
		keys = append(keys, key)
	}
	return
}
