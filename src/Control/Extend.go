func ArrayExtend(f func([]any) any, arr []any) []any {
	res := make([]any, len(arr))
	for i := range arr {
		res[i] = f(arr[i:])
	}
	return res
}
