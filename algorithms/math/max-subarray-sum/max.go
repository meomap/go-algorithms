// Package max_subarray provides Max Subarray Sum implementation
package max_subarray

import (
	"fmt"
	"math"
)

func BruteForce(in []int) int {
	var max int
	lenIn := len(in)
	for i := 0; i < lenIn; i++ {
		currentSum := in[i]
		for j := i + 1; j < lenIn; j++ {
			currentSum += in[j]
			if currentSum > max {
				max = currentSum
			}
		}
	}
	return max
}

// https://www.geeksforgeeks.org/divide-and-conquer-maximum-sum-subarray/
func DivideAndConquor(in []int) int {
	return maxSubArray(in, 0, len(in)-1)
}

// https://en.wikipedia.org/wiki/Maximum_subarray_problem
// https://www.geeksforgeeks.org/largest-sum-contiguous-subarray/
// Complexity: O(n)
func DynamicProgramming(in []int) int {
	lenIn := len(in)
	maxSoFar, maxEndingHere := in[0], 0
	for i := 1; i < lenIn; i++ {
		// either in[i] is a prefix of next maxEnding or not
		maxEndingHere = int(math.Max(float64(in[i]), float64(maxEndingHere+in[i])))
		maxSoFar = int(math.Max(float64(maxSoFar), float64(maxEndingHere)))
	}
	return maxSoFar
}

func maxSubArray(in []int, lo, hi int) int {
	// basecase
	lenIn := hi - lo
	if lenIn == 1 {
		return in[lo]
	} else if lenIn < 1 {
		return 0
	}
	mid := (hi + lo) / 2

	// divide
	maxL := maxSubArray(in, lo, mid)
	maxR := maxSubArray(in, mid, hi)

	// combine
	maxCross := maxCrossSubArray(in, lo, mid, hi)
	if maxCross == 8 {
		fmt.Printf("maxCross = %+v\n", maxCross)
	}
	return int(math.Max(math.Max(float64(maxL), float64(maxR)), float64(maxCross)))
}

func maxCrossSubArray(in []int, lo, mid, hi int) int {
	// join max found from middle to left & middle to right
	maxL, sum := 0, 0
	for i := mid; i >= lo; i-- {
		sum += in[i]
		if sum > maxL {
			maxL = sum
		}
	}

	maxR, sum := 0, 0
	for i := mid + 1; i <= hi; i++ {
		sum += in[i]
		if sum > maxR {
			maxR = sum
		}
	}
	return maxL + maxR
}
