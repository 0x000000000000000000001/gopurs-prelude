func ArrayMap(f func(any) any, arr []any) []any {
	result := make([]any, len(arr))
	for i, v := range arr {
		result[i] = f(v)
	}
	return result
}
