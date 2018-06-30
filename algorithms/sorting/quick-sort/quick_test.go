// Package quick provides QuickSort implementation
package quick

import (
	"fmt"
	"testing"

	"github.com/meomap/go-algorithms/utils"
)

func TestQuickSort(t *testing.T) {
	lst := utils.RandNumbers(6)

	sorted := Sort(lst)
	utils.TestListSorted(sorted, t)
	fmt.Printf("lst = %+v\n", lst)
	fmt.Printf("sorted = %+v\n", sorted)
}
