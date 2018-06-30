// Package merge provides Merge Sort Algorithm
package merge

import (
	"fmt"
	"testing"

	"github.com/meomap/go-algorithms/utils"
)

var compareFn = func(a, b interface{}) int {
	numA, ok := a.(uint)
	if !ok {
		panic("Expect a to be uint number")
	}
	numB, ok := b.(uint)
	if !ok {
		panic("Expect b to be uint number")
	}

	if numA > numB {
		return 1
	} else if numA == numB {
		return 0
	}
	return -1
}

func randNumbers(n uint) []interface{} {
	numbers := utils.RandNumbers(n)
	lst := make([]interface{}, n)
	for i := uint(0); i < n; i++ {
		lst[i] = numbers[i]
	}
	return lst
}

func TestMergeSort(t *testing.T) {
	lst := randNumbers(10)
	sorted := Sort(lst, compareFn)
	utils.TestListSorted(sorted, t)
	fmt.Printf("lst = %+v\n", lst)
	fmt.Printf("sorted = %+v\n", sorted)
}

func BenchmarkMergeSort1000(b *testing.B) {
	lst := randNumbers(1000)
	for i := 0; i < b.N; i++ {
		Sort(lst, compareFn)
	}
}

func BenchmarkMergeSort10000(b *testing.B) {
	lst := randNumbers(10000)
	for i := 0; i < b.N; i++ {
		Sort(lst, compareFn)
	}
}

func BenchmarkMergeSort100000(b *testing.B) {
	lst := randNumbers(100000)
	for i := 0; i < b.N; i++ {
		Sort(lst, compareFn)
	}
}
