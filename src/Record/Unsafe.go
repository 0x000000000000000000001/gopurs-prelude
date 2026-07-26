func UnsafeHas(label string, rec map[string]interface{}) bool {
	_, ok := rec[label]
	return ok
}
func UnsafeGet(label string, rec map[string]interface{}) interface{} {
	return rec[label]
}
func UnsafeSet(label string, value interface{}, rec map[string]interface{}) map[string]interface{} {
	newMap := make(map[string]interface{}, len(rec)+1)
	for k, v := range rec {
		newMap[k] = v
	}
	newMap[label] = value
	return newMap
}
func UnsafeDelete(label string, rec map[string]interface{}) map[string]interface{} {
	newMap := make(map[string]interface{})
	for k, v := range rec {
		if k != label {
			newMap[k] = v
		}
	}
	return newMap
}
