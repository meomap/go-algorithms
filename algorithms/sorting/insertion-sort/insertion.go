// Package insertion provides Insertion Sort Implementation
package insertion

func Sort(in []uint) (out []uint) {
	numLen := len(in)
	out = in
	for i := 1; i < numLen; i++ {
		x, j := out[i], i-1
		for ; j > 0 && x < out[j]; j-- {
			// move all previous position forward
			// until find correct position of x
			out[j+1] = out[j]
		}
		// return x to right place
		out[j] = x
	}
	return
}
