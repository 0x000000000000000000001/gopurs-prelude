func ArrayApply(fs []func(any) any, xs []any) []any {
	result := make([]any, 0, len(fs)*len(xs))
	for _, f := range fs {
		for _, x := range xs {
			result = append(result, f(x))
		}
	}
	return result
}
