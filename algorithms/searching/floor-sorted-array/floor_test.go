// Package floor provides Floor in a Sorted Array test
package floor

import (
	"math/rand"
	"testing"
	"time"
)

type testCase struct {
	in    []int
	x     int
	floor int
}

var data = []testCase{
	{
		in:    []int{1, 2, 8, 10, 10, 12, 19},
		x:     5,
		floor: 2,
	},
	{
		in:    []int{1, 2, 8, 10, 10, 12, 19},
		x:     20,
		floor: 19,
	},
	{
		in:    []int{1, 2, 8, 10, 10, 12, 19},
		x:     0,
		floor: -1,
	},
}

func TestLinearScan(t *testing.T) {
	for _, c := range data {
		out := searchLinear(c.in, c.x)
		if c.floor != out {
			t.Errorf("Expect %d but got %d", c.floor, out)
		}
	}
}

func TestBinaryScan(t *testing.T) {
	for _, c := range data {
		out := searchBinary(c.in, c.x)
		if c.floor != out {
			t.Errorf("Expect %d but got %d", c.floor, out)
		}
	}
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randArrays(length, min, max int) []int {
	out := make([]int, length)
	for i := 0; i < length; i++ {
		out[i] = randInt(min, max)
	}
	return out
}

func testSet(count int) ([]int, int) {
	floor := rand.Intn(count)
	lenLeft := rand.Intn(count)
	return append(randArrays(lenLeft, 0, floor), randArrays(count-lenLeft, floor, count)...), floor
}

func BenchmarkLinear1000(b *testing.B) {
	in, x := testSet(1000)
	for i := 0; i < b.N; i++ {
		searchLinear(in, x)
	}
}

func BenchmarkBinary1000(b *testing.B) {
	in, x := testSet(1000)
	for i := 0; i < b.N; i++ {
		searchBinary(in, x)
	}
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
