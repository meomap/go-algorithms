// Package closest provides ClosestPair Implementation
package closest

import "testing"

var (
	sampleData1       = []Point{{2, 3}, {12, 30}, {40, 50}, {5, 1}, {12, 10}, {3, 4}}
	sampleExpectPair1 = []Point{{2, 3}, {3, 4}}
	// sample for falling into split points
	sampleData       = []Point{{1, 10}, {2, 20}, {3, 4}, {4, 5}, {5, 30}, {6, 70}}
	sampleExpectPair = []Point{{3, 4}, {4, 5}}
)

func TestBruteForce(t *testing.T) {
	points := make([]Point, 10)
	// create a random set of points with expected closest pair is (0,0), (1,1)
	for i := 9; i >= 0; i-- {
		a := NewPoint(float64(i), float64(i*i))
		points[i] = a
	}
	a, b := points[0], points[1]
	pair := BruteForce(points)
	if pair.Start != a || pair.End != b {
		t.Errorf("Expect closest pair matching (0,0) (1,1) but got %q", pair)
	}
}

func TestBruteForceSample(t *testing.T) {
	pair := BruteForce(sampleData)
	if pair.Start != sampleExpectPair[0] || pair.End != sampleExpectPair[1] {
		t.Errorf("Expect closest pair matching %q %q but got %q", sampleExpectPair[0], sampleExpectPair[1], pair)
	}
}

func TestDivideAndConquor(t *testing.T) {
	pair := DivideAndConquor(sampleData)
	if pair.Start != sampleExpectPair[0] || pair.End != sampleExpectPair[1] {
		t.Errorf("Expect closest pair matching %q %q but got %q", sampleExpectPair[0], sampleExpectPair[1], pair)
	}
}

func BenchmarkBruteForce100(b *testing.B) {
	times := 100
	points := make([]Point, times)
	for i := times - 1; i >= 0; i-- {
		a := NewPoint(float64(i), float64(i*i))
		points[i] = a
	}

	for i := 0; i < b.N; i++ {
		BruteForce(points)
	}
}

func BenchmarkDivideAndConquor100(b *testing.B) {
	times := 100
	points := make([]Point, times)
	for i := times - 1; i >= 0; i-- {
		a := NewPoint(float64(i), float64(i*i))
		points[i] = a
	}

	for i := 0; i < b.N; i++ {
		DivideAndConquor(points)
	}
}
