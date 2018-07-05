package longest

import "testing"

type testData struct {
	in       []string
	expected string
}

var testCases = []testData{
	{
		in:       []string{"geeksforgeeks", "geeks", "geek", "geezer"},
		expected: "gee",
	},
	{
		in:       []string{"apple", "ape", "april"},
		expected: "ap",
	},
}

func TestBruteForce(t *testing.T) {
	for _, c := range testCases {
		out, err := BruteForce(c.in)
		if err != nil {
			t.Fatalf("Unexpected error %q", err)
		}
		if c.expected != out {
			t.Errorf("Expect %q but got %q", c.expected, out)
		}
	}
}

func TestDivideAndConquor(t *testing.T) {
	for _, c := range testCases {
		out, err := DivideAndConquor(c.in)
		if err != nil {
			t.Fatalf("Unexpected error %q", err)
		}
		if c.expected != out {
			t.Errorf("Expect %q but got %q", c.expected, out)
		}
	}
}
