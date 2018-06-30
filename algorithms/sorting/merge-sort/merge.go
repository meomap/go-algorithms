// Package merge provides Merge Sort Implementation
package merge

type CompareFunc func(a, b interface{}) int

// Sort implement Merge Sort algorithm
func Sort(in []interface{}, compareFn CompareFunc) (out []interface{}) {
	out = in
	numLen := len(out)
	if numLen == 1 {
		// base case
		return
	}
	med := numLen / 2
	// divide
	left, right := out[:med], out[med:]

	// recursive sort
	left, right = Sort(left, compareFn), Sort(right, compareFn)

	// merge
	out = merge(left, right, compareFn)
	return
}

func merge(left, right []interface{}, compareFn CompareFunc) (out []interface{}) {
	var i, j int
	for {
		if compareFn(left[i], right[j]) == -1 {
			out = append(out, left[i])
			i++
		} else {
			out = append(out, right[j])
			j++
		}
		if i == len(left) || j == len(right) {
			// either left or right exhausted
			break
		}
	}

	// push all remaining
	if i == len(left) {
		out = append(out, right[j:]...)
	} else {
		out = append(out, left[i:]...)
	}
	return
}
