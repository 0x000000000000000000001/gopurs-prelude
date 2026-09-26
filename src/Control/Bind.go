// ConcatMap natif sur les tableaux PureScript : entree et sortie en `Value`
// (aucune copie de conversion), le callback est applique element par element.
func ArrayBind(arr []gopurs_runtime.Value, f func(gopurs_runtime.Value) []gopurs_runtime.Value) gopurs_runtime.Value {
	result := make([]gopurs_runtime.Value, 0, len(arr))
	for _, v := range arr {
		result = append(result, f(v)...)
	}
	return gopurs_runtime.Array(result)
}
