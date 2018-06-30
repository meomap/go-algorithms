package utils

// TestingT is an interface wrapper around *testing.T
type testingT interface {
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

func TestListSorted(lst []interface{}, t testingT) {
	for i := 0; i < len(lst)-2; i++ {
		current, _ := lst[i].(uint)
		next, _ := lst[i+1].(uint)
		if current > next {
			t.Fatalf("Index %d got current=%d next=%d", i, current, next)
		}
	}
}
