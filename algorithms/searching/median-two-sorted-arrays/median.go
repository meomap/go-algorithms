// Package median provides Median Of Two Sorted Arrays implementation
package median

import (
	"errors"
	"fmt"
)

func validateSorted(in []int) error {
	if len(in) == 1 {
		return nil
	}
	for i := 1; i < len(in); i++ {
		if in[i-1] > in[i] {
			return errors.New(fmt.Sprintf("Array not sorted, index[%d]=%d index[%d]=%d", i-1, in[i-1], i, in[i]))
		}
	}
	return nil
}

// SearchByMerging merges two sorted array then find the median by average value
// Complexity: N
func SearchByMerging(l, r []int) int {
	lenL, lenR := len(l), len(r)
	i, j := 0, 0
	merged := []int{}
	for i < lenL && j < lenR {
		if l[i] < r[j] {
			merged = append(merged, l[i])
			i++
		} else {
			merged = append(merged, r[j])
			j++
		}
	}
	if i == lenL {
		// left exhausted, add remaining from right
		merged = append(merged, r[j:]...)
	} else {
		merged = append(merged, l[i:]...)
	}
	lenM := lenL + lenR
	// two array of size 2n. after merging pick two middle elements
	mid1 := (lenM - 2) / 2
	mid2 := lenM / 2
	return (merged[mid1] + merged[mid2]) / 2
}
