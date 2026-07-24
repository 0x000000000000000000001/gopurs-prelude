func ArrayBind(arr []any, f func(any) []any) []any {
	var result []any
	for _, v := range arr {
		result = append(result, f(v)...)
	}
	return result
}
