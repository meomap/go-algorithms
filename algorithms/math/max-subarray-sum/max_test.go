// Package max_subarray provides Max Subarray Sum implementation
package max_subarray

import "testing"

var (
	sample    = []int{-2, -5, 6, -2, -3, 1, 5, -6}
	expectMax = 7
)

func TestBruteForce(t *testing.T) {
	out := BruteForce(sample)
	if out != expectMax {
		t.Errorf("Expect max value %d but got %d", expectMax, out)
	}
}

func TestDivideAndConquor(t *testing.T) {
	out := DivideAndConquor(sample)
	if out != expectMax {
		t.Errorf("Expect max value %d but got %d", expectMax, out)
	}
}

func TestDynamicProgramming(t *testing.T) {
	out := DynamicProgramming(sample)
	if out != expectMax {
		t.Errorf("Expect max value %d but got %d", expectMax, out)
	}
}

// generate sample set by prefixing & sufficing origin sample set data
// with negative numbers
func createSampleSet(numbers int) []int {
	lenOrigin := len(sample)
	lenPrefixL := (numbers - lenOrigin) / 2
	lenPrefixR := numbers - lenOrigin - lenPrefixL

	largeSet := make([]int, numbers)
	for i := 0; i < lenPrefixL; i++ {
		largeSet[i] = -1 * (i + 1)
	}
	for i := 0; i < lenOrigin; i++ {
		largeSet[i] = sample[i]
	}
	for i := 0; i < lenPrefixR; i++ {
		largeSet[i] = -1 * (i + 1)
	}
	return largeSet
}

func BenchmarkBruteForce1000(b *testing.B) {
	largeSet := createSampleSet(1000)
	for i := 0; i < b.N; i++ {
		BruteForce(largeSet)
	}
}

func BenchmarkDivideAndConquor1000(b *testing.B) {
	largeSet := createSampleSet(1000)
	for i := 0; i < b.N; i++ {
		DivideAndConquor(largeSet)
	}
}

func BenchmarkDynamicProgramming1000(b *testing.B) {
	largeSet := createSampleSet(1000)
	for i := 0; i < b.N; i++ {
		DynamicProgramming(largeSet)
	}
}
