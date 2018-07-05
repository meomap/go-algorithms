// Package median provides Median Of Two Sorted Arrays implementation
// https://www.geeksforgeeks.org/median-of-two-sorted-arrays/
package median

import "math"

// SearchByLinear merges two sorted array first then find the median by average value
// from two middle value
// Complexity: N
func SearchByLinear(l, r []int) int {
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

// return input array median value
// i.e {15, 26, 38} => 26
func median(in []int) (int, int) {
	if len(in) == 0 {
		return -1, -1
	}
	idx := len(in) / 2
	return idx, in[idx]
}

// SearchByComparingMedians compares the medians found from two subsets then
// reduce subsets size until reasonable solution found
func SearchByComparingMedians(l, r []int) int {
	lenL, lenR := len(l), len(r)
	if lenL == lenR && lenL == 2 {
		// [l1:l2] [r1:r2] => probably median is between l2, r1
		return int(math.Max(float64(l[0]), float64(r[0]))+math.Min(float64(l[1]), float64(r[1]))) / 2
	}
	lidx, lmed := median(l)
	ridx, rmed := median(r)
	if lmed == rmed {
		// found matching median
		return lmed
	}
	if lmed < rmed {
		// median is at either at left[lm:] or right[:rm+1]
		return SearchByComparingMedians(l[lidx:], r[:ridx+1])
	}
	return SearchByComparingMedians(l[:lidx+1], r[ridx:])
}
