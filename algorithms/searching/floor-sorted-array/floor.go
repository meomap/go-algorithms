// Package floor provides Floor in a Sorted Array implementation
// https://www.geeksforgeeks.org/floor-in-a-sorted-array/
package floor

// given a sorted array and a value x, the floor of x is the largest element
// in array smaller than or equal to
// Complexity: N
func searchLinear(in []int, x int) int {
	lenIn := len(in)
	if in[0] > x {
		return -1
	}
	for i := 1; i < lenIn; i++ {
		if in[i] > x {
			return in[i-1]
		}
	}
	return in[lenIn-1]
}

func searchBinary(in []int, x int) int {
	lenIn := len(in)
	if lenIn == 1 {
		if in[0] > x {
			return -1
		}
		return in[0]
	}
	mid := lenIn / 2
	if in[mid] > x {
		// floor value must be at far left of array
		return searchBinary(in[:mid], x)
	}
	return searchBinary(in[mid:], x)
}
