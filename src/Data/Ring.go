import "gopurs/output/gopurs_runtime"

func IntSub(x int64, y int64) int64 {
	return gopurs_runtime.IntSub(x, y)
}
func NumSub(x float64, y float64) float64 {
	return x - y
}
