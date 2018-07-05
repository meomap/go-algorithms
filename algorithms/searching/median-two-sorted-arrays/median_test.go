// Package median provides Median Of Two Sorted Arrays test
package median

import "testing"

var (
	in1      = []int{1, 12, 15, 26, 38}
	in2      = []int{2, 13, 17, 30, 45}
	expected = 16
)

func TestNSolution(t *testing.T) {
	out := SearchByMerging(in1, in2)
	if expected != out {
		t.Errorf("Expect %d but got %d", expected, out)
	}
}
