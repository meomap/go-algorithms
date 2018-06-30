// Package bubble provides bubble sort implementation
package bubble

func Sort(in []uint) (out []uint) {
	numLen := len(in)
	out = in
	for i := 0; i < numLen; i++ {
		swap := false
		for j := numLen - 1; j >= i+1; j-- {
			if out[j] < out[j-1] {
				out[j], out[j-1] = out[j-1], out[j]
				swap = true
			}
		}
		if !swap {
			// already sorted now
			break
		}
	}
	return
}
