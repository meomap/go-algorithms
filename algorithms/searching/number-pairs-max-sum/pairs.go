// Package pairs provides Number of Pairs with Max sum implementation
// https://www.geeksforgeeks.org/number-pairs-maximum-sum/
package pairs

// Complexity: O(n^2)
func bruteForce(in []int) int {
	lenIn := len(in)
	if lenIn <= 1 {
		return 0
	}
	count, max := 1, 0
	for i := 0; i < lenIn; i++ {
		for j := i + 1; j < lenIn; j++ {
			sum := in[i] + in[j]
			if sum > max {
				// new max value
				count = 1
				max = sum
			} else if sum == max {
				count++
			}
		}
	}
	return count
}
