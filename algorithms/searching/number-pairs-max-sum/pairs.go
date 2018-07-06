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

// find maxium sum & maxium number value then do linear scan.
// note that maxium number in array should fall into maximum sum
// Complexity: O(n)
func linearScan(in []int) int {
	lenIn := len(in)
	i, j := 0, lenIn-1
	sumMax, numMax := 0, 0
	for i < j {
		if in[i] < in[j] {
			if in[j] > numMax {
				numMax = in[j]
			}
			i++
		} else {
			if in[i] > numMax {
				numMax = in[i]
			}
			j--
		}
		sum := in[i] + in[j]
		if sum > sumMax {
			sumMax = sum
		}
	}
	count := 0

	for i := 0; i < lenIn; i++ {
		if in[i] == numMax || in[i] == sumMax-numMax {
			count++
		}
	}
	return count
}
