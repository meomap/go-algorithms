// Package insertion provides Insertion sort
// https://github.com/TheAlgorithms/Go#insertion

package insertion

import (
	"fmt"
	"testing"

	"github.com/meomap/go-algorithms/utils"
)

func TestInsertionSort(t *testing.T) {
	lst := utils.RandNumbers(10)

	sorted := Sort(lst)
	utils.TestListSorted(sorted, t)
	fmt.Printf("lst = %+v\n", lst)
	fmt.Printf("sorted = %+v\n", sorted)
}

func BenchmarkInsertionSort1000(b *testing.B) {
	lst := utils.RandNumbers(1000)
	for i := 0; i < b.N; i++ {
		Sort(lst)
	}
}

func BenchmarkInsertionSort10000(b *testing.B) {
	lst := utils.RandNumbers(10000)
	for i := 0; i < b.N; i++ {
		Sort(lst)
	}
}

func BenchmarkInsertionSort100000(b *testing.B) {
	lst := utils.RandNumbers(100000)
	for i := 0; i < b.N; i++ {
		Sort(lst)
	}
}
