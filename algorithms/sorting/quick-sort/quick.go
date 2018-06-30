// Package quick provides QuickSort implementation
// https://en.wikipedia.org/wiki/Quicksort
package quick

import "fmt"

// Sort implement QuickSort algorithm
func Sort(in []uint) (out []uint) {
	fmt.Printf("in = %+v\n", in)
	out = in
	quickSort(out, 0, len(in)-1)
	return
}

// sort inplace from lo to high index
func quickSort(lst []uint, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(lst, lo, hi)
	quickSort(lst, lo, p-1)
	quickSort(lst, p+1, hi)
}

// return pivot position which partions list into two halves
func partition(lst []uint, lo, hi int) int {
	// for easy of implementation, choose last value as pivot
	pivot := lst[hi]
	// index for pivot pre-correct position, also Top of left partition with smaller values
	// start by negative one
	i := lo - 1

	for j := lo; j <= hi-1; j++ {
		if lst[j] < pivot {
			// push to left partion
			i++
			lst[j], lst[i] = lst[i], lst[j]
		}
	}
	// move pivot to correct position
	i++
	lst[i], lst[hi] = pivot, lst[i]
	return i
}
