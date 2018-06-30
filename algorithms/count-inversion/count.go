package count

// CountInversion implement counting inversion algorithm
func CountInversion(in []int) int {
	_, cnt := sortAndCount(in)
	return cnt
}

// CountInversion implement algorithm to count array inversion
// https://www.geeksforgeeks.org/counting-inversions/
func sortAndCount(in []int) ([]int, int) {
	if len(in) == 1 {
		// base case
		return in, 0
	}
	mid := len(in) / 2
	lsorted, lc := sortAndCount(in[:mid])
	rsorted, rc := sortAndCount(in[mid:])
	sorted, mc := mergeAndCount(lsorted, rsorted)
	return sorted, lc + rc + mc
}

// mergeAndCount returns sorted array together with inversion count
func mergeAndCount(lsorted, rsorted []int) ([]int, int) {
	i, j := 0, 0
	sorted, count := []int{}, 0
	for {
		if lsorted[i] <= rsorted[j] {
			sorted = append(sorted, lsorted[i])
			i++
		} else {
			sorted = append(sorted, rsorted[j])
			j++
			// a[i] > b[j] thus there's inversion in range a[i:]
			count += len(lsorted) - i
		}
		if i == len(lsorted) || j == len(rsorted) {
			break
		}
	}

	if i == len(lsorted) {
		// left exhausted, collect right
		sorted = append(sorted, rsorted...)
	} else {
		sorted = append(sorted, lsorted...)
	}
	return sorted, count
}
