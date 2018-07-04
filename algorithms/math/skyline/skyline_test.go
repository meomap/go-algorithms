// Package skyline provides Skyline Problem tests
package skyline

import "testing"

var multiple = testCase{
	in: []building{
		{1, 11, 5}, {2, 6, 7}, {3, 13, 9}, {12, 7, 16}, {14, 3, 25},
		{19, 18, 22},
	},
	expect: []strip{
		{1, 11}, {3, 13}, {9, 0}, {12, 7}, {16, 3}, {19, 18},
		{22, 3}, {25, 0},
	},
}

type testCase struct {
	in     []building
	expect []strip
}

func TestSkylineBruteforce(t *testing.T) {
	// simple case - 1 building
	in1 := []building{
		{1, 11, 5},
	}
	expect1 := []strip{
		{1, 11}, {5, 0},
	}
	out1 := BruteForce(in1)
	assertStripsMatched(t, expect1, out1)

	// multiple buildings
	out2 := BruteForce(multiple.in)
	assertStripsMatched(t, multiple.expect, out2)
}

func TestSkylineDivideAndConquor(t *testing.T) {
	out := DivideAndConquor(multiple.in)
	assertStripsMatched(t, multiple.expect, out)
}

func assertStripsMatched(t *testing.T, expected, actual []strip) {
	//fmt.Printf("expected=%+v\n actual=%+v \n", expected, actual)
	lenE, lenA := len(expected), len(actual)
	if lenE != lenA {
		t.Errorf("Expect length %d but got length %d\n", lenE, lenA)
	}
	for i := 0; i < lenA; i++ {
		if expected[i] != actual[i] {
			t.Errorf("Index %d: expect %+v but got %+v", i, expected[i], actual[i])
		}
	}
}
