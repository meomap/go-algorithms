package longest

import "testing"

var (
	testIn       = []string{"geeksforgeeks", "geeks", "geek", "geezer"}
	testExpected = "gee"
)

func TestBruteForce(t *testing.T) {
	out, err := BruteForce(testIn)
	if err != nil {
		t.Fatalf("Unexpected error %q", err)
	}
	if testExpected != out {
		t.Errorf("Expect %q but got %q", testExpected, out)
	}
}

func TestDivideAndConquor(t *testing.T) {

}
