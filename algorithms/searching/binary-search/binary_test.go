// Package binary provides BinarySearch implementation
package binary

import (
	"fmt"
	"testing"

	merge "github.com/meomap/go-algorithms/algorithms/sorting/merge-sort"
	"github.com/meomap/go-algorithms/utils"
)

func TestSearch(t *testing.T) {
	lst := utils.RandNumbers(10)
	sorted := merge.Sort(lst)

	index := Search(sorted, sorted[8])
	if index == -1 {
		t.Errorf("Expect got result for value %d, in=%+v", sorted[8], sorted)
	}
	fmt.Printf("index= %+v\n", index)
	fmt.Printf("sorted = %+v\n", sorted)
	fmt.Printf("sorted[8] = %+v\n", sorted[8])
	index = Search(sorted, 9999)
	if index != -1 {
		t.Errorf("Expect got result -1 for value %d, in=%+v", 9999, sorted)
	}
}
