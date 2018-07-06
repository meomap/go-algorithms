// Package closest provides Closet Number in Array implementation
package closest

import "math"

func searchLinear(in []int, target int) int {
	lenIn := len(in)
	if lenIn == 1 {
		return in[0]
	}
	i := 0
	for {
		i++
		if i == lenIn-1 || in[i] >= target {
			break
		}
	}
	n1, n2 := in[i], in[i-1]
	d1 := math.Abs(float64(n1 - target))
	d2 := math.Abs(float64(n2 - target))
	if d1 < d2 {
		return n1
	}
	return n2
}
