package count

import "testing"

func TestCountingInversion(t *testing.T) {
	lst := []int{2, 4, 1, 3, 5}
	expect := 3
	cnt := CountInversion(lst)
	if cnt != expect {
		t.Fatalf("Expect %d but got %d", expect, cnt)
	}
}
