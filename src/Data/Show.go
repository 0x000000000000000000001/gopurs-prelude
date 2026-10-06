import (
	"fmt"
	"math"
	"strconv"
	"strings"
)
func ShowIntImpl(n int64) string {
	return fmt.Sprintf("%v", n)
}
func ShowNumberImpl(n float64) string {
	if n == 0 {
		// JavaScript Number#toString renders both signed zeros as "0".
		return "0.0"
	} else if math.IsNaN(n) {
		return "NaN"
	} else if math.IsInf(n, 1) {
		return "Infinity"
	} else if math.IsInf(n, -1) {
		return "-Infinity"
	}

	absN := math.Abs(n)
	if absN >= 1e21 || absN < 1e-6 {
		// Keep the shortest round-tripping digits, with JS notation thresholds
		// and an unpadded exponent (Go's e-07 becomes JavaScript's e-7).
		str := strconv.FormatFloat(n, 'e', -1, 64)
		exponent := strings.IndexByte(str, 'e')
		return str[:exponent+2] + strings.TrimLeft(str[exponent+2:], "0")
	}

	str := strconv.FormatFloat(n, 'f', -1, 64)
	if strings.Contains(str, ".") {
		return str
	}
	return str + ".0"
}
func ShowCharImpl(c string) string {
	return fmt.Sprintf("'%s'", c)
}
func ShowStringImpl(s string) string {
	return fmt.Sprintf("%q", s)
}
func ShowArrayImpl(f func(interface{}) string, arr []interface{}) string {
	res := "["
	for i, v := range arr {
		if i > 0 {
			res += ","
		}
		res += f(v)
	}
	res += "]"
	return res
}
