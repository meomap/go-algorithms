// Package karatsuba provides Binary String Multiplication impelementation
package karatsuba

import "testing"

func TestBruteForce(t *testing.T) {
	a, b := "1100", "1010"
	expected := 120
	out := BruteForce([]byte(a), []byte(b))
	if out != expected {
		t.Errorf("Expect %d but got %d", expected, out)
	}
}
