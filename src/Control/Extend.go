func ArrayExtend(f func([]interface{}) interface{}, arr []interface{}) []interface{} {
	res := make([]interface{}, len(arr))
	for i := range arr {
		res[i] = f(arr[i:])
	}
	return res
}
