// Package binary provides BinarySearch implementation
package binary

func Search(lst []uint, value uint) int {
	return binarySearch(lst, 0, len(lst), value)
}

func binarySearch(lst []uint, lo, hi int, value uint) int {
	mid := (hi + lo) / 2
	if lst[mid] == value {
		return mid
	} else if hi-lo <= 1 {
		return -1
	} else if lst[mid] < value {
		return binarySearch(lst, mid, hi, value)
	} else {
		return binarySearch(lst, lo, mid, value)
	}
}
