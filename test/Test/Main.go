package Test_Main

import (
	"fmt"
	"math"
	"gopurs/output/gopurs_runtime"
)

func TestNumberShow(showNumber func(interface{}) interface{}) func(interface{}) interface{} {
	return func(interface{}) interface{} {
		cases := []struct {
			num      float64
			expected string
		}{
			// Within Int range
			{0.0, "0.0"},
			{1.0, "1.0"},
			{-1.0, "-1.0"},
			{500.0, "500.0"},

			// Outside Int range
			{1e10, "10000000000.0"},
			{1e10 + 0.5, "10000000000.5"},
			{-1e10, "-10000000000.0"},
			{-1e10 - 0.5, "-10000000000.5"},

			// With exponent
			{1e21, "1e+21"},
			{1e-21, "1e-21"},

			// With decimal and exponent
			{1.5e21, "1.5e+21"},
			{1.5e-10, "1.5e-10"},

			{math.NaN(), "NaN"},
			{math.Inf(1), "Infinity"},
			{math.Inf(-1), "-Infinity"},
		}

		for _, c := range cases {
			actualVal := showNumber(c.num)
			actual := gopurs_runtime.Unbox[string](actualVal)
			if actual != c.expected {
				panic(fmt.Sprintf("For %f, expected %s, got: %s.", c.num, c.expected, actual))
			}
		}
		return nil
	}
}

func MakeArray(length int) []interface{} {
	arr := make([]interface{}, length)
	for i := 0; i < length; i++ {
		arr[i] = i
	}
	return arr
}
