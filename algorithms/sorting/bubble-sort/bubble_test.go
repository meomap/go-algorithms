// Package bubble provides BubbleSort implementation
package bubble

import (
	"testing"

	"github.com/meomap/go-algorithms/utils"
)

func TestBubbleSort(t *testing.T) {
	lst := utils.RandNumbers(100)

	sorted := Sort(lst)

}

func BenchmarkBubbleSort1000(b *testing.B) {
	lst := utils.RandNumbers(1000)
	for i := 0; i < b.N; i++ {
		Sort(lst)
	}
}

func BenchmarkBubbleSort10000(b *testing.B) {
	lst := utils.RandNumbers(10000)
	for i := 0; i < b.N; i++ {
		Sort(lst)
	}
}

func BenchmarkBubbleSort100000(b *testing.B) {
	lst := utils.RandNumbers(100000)
	for i := 0; i < b.N; i++ {
		Sort(lst)
	}
}
