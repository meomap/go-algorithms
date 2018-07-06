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
	return minDist(in[i], in[i-1], target)
}

func minDist(n1, n2, target int) int {
	d1 := math.Abs(float64(n1 - target))
	d2 := math.Abs(float64(n2 - target))
	if d1 < d2 {
		return n1
	}
	return n2
}

func searchBinary(in []int, target int) int {
	lenIn := len(in)
	if lenIn == 1 {
		return in[0]
	} else if lenIn == 2 {
		return minDist(in[0], in[1], target)
	}
	mid := lenIn / 2
	if in[mid] > target {
		// expand to left
		return searchBinary(in[:mid+1], target)
	}
	return searchBinary(in[mid:], target)
}
