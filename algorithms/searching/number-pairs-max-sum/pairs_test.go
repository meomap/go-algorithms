// Package  provides Number of Pairs with Max sum test
package pairs

import "testing"

type testCase struct {
	in    []int
	count int
}

var data = []testCase{
	{
		in:    []int{1, 1, 1, 2, 2, 2},
		count: 3,
	},
	{
		in:    []int{1, 4, 3, 3, 5, 1},
		count: 1,
	},
}

type searchAlgo func(i []int) int

func runTest(t *testing.T, testData []testCase, algorithm searchAlgo) {
	for _, c := range testData {
		out := algorithm(c.in)
		if c.count != out {
			t.Errorf("Expect %d but got %d", c.count, out)
		}
	}
}

func TestBruteForce(t *testing.T) {
	runTest(t, data, bruteForce)
}
