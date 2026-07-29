package Record_Unsafe

import "gopurs/output/gopurs_runtime"

func UnsafeHas(label string, recVal interface{}) bool {
	rec := gopurs_runtime.RecordToMap(recVal.(gopurs_runtime.Value))
	_, ok := rec[label]
	return ok
}
func UnsafeGet(label string, recVal interface{}) interface{} {
	rec := gopurs_runtime.RecordToMap(recVal.(gopurs_runtime.Value))
	return rec[label]
}
func UnsafeSet(label string, value interface{}, recVal interface{}) interface{} {
	rec := gopurs_runtime.RecordToMap(recVal.(gopurs_runtime.Value))
	newMap := make(map[string]gopurs_runtime.Value, len(rec)+1)
	for k, v := range rec {
		newMap[k] = v
	}
	newMap[label] = gopurs_runtime.Any(value)
	return gopurs_runtime.Record(newMap)
}
func UnsafeDelete(label string, recVal interface{}) interface{} {
	rec := gopurs_runtime.RecordToMap(recVal.(gopurs_runtime.Value))
	newMap := make(map[string]gopurs_runtime.Value)
	for k, v := range rec {
		if k != label {
			newMap[k] = v
		}
	}
	return gopurs_runtime.Record(newMap)
}
