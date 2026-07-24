func UnsafeHas(label string, rec map[string]any) bool {
	_, ok := rec[label]
	return ok
}
func UnsafeGet(label string, rec map[string]any) any {
	return rec[label]
}
func UnsafeSet(label string, value any, rec map[string]any) map[string]any {
	newMap := make(map[string]any, len(rec)+1)
	for k, v := range rec {
		newMap[k] = v
	}
	newMap[label] = value
	return newMap
}
func UnsafeDelete(label string, rec map[string]any) map[string]any {
	newMap := make(map[string]any)
	for k, v := range rec {
		if k != label {
			newMap[k] = v
		}
	}
	return newMap
}
