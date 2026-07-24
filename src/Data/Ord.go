func OrdBooleanImpl(lt any, eq any, gt any, x bool, y bool) any {
	if !x && y {
		return lt
	} else if x == y {
		return eq
	}
	return gt
}
func OrdIntImpl(lt any, eq any, gt any, x int, y int) any {
	if x < y { return lt }
	if x == y { return eq }
	return gt
}
func OrdCharImpl(lt any, eq any, gt any, x string, y string) any {
	if x < y { return lt }
	if x == y { return eq }
	return gt
}
func OrdStringImpl(lt any, eq any, gt any, x string, y string) any {
	if x < y { return lt }
	if x == y { return eq }
	return gt
}
func OrdNumberImpl(lt any, eq any, gt any, x float64, y float64) any {
	if x < y { return lt }
	if x == y { return eq }
	return gt
}
func OrdArrayImpl(f func(any) func(any) int, xs []any, ys []any) int {
	xlen := len(xs)
	ylen := len(ys)
	for i := 0; i < xlen && i < ylen; i++ {
		o := f(xs[i])(ys[i])
		if o != 0 {
			return o
		}
	}
	if xlen == ylen {
		return 0
	} else if xlen > ylen {
		return 1
	} else {
		return -1
	}
}
