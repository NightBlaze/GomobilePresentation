package array

func Map[T any, U any](array []T, f func(T) U) []U {
	mappedArray := make([]U, len(array))
	for i, value := range array {
		mappedArray[i] = f(value)
	}
	return mappedArray
}
