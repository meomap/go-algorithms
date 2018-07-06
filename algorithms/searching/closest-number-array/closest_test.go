// Package closest provides Closet Number in Array test
package closest

import "testing"

type testCase struct {
	in      []int
	target  int
	closest int
}

var data = []testCase{
	{
		in:      []int{1, 2, 4, 5, 6, 6, 8, 9},
		target:  11,
		closest: 9,
	},
	{
		in:      []int{2, 5, 6, 7, 8, 8, 9},
		target:  4,
		closest: 5,
	},
}

type searchAlgo func([]int, int) int

func runTest(t *testing.T, testData []testCase, algorithm searchAlgo) {
	for _, c := range testData {
		out := algorithm(c.in, c.target)
		if c.closest != out {
			t.Errorf("Expect %d but got %d", c.closest, out)
		}
	}
}

func TestLinear(t *testing.T) {
	runTest(t, data, searchLinear)
}
