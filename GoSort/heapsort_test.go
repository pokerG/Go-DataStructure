package GoSort

import (
	"testing"
)

func BenchmarkHeap1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := testDataSorted()
		HeapSort(input)
	}
}

func BenchmarkHeap2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := testDataReversed()
		HeapSort(input)
	}
}

func BenchmarkHeap3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := testDataRandom()
		HeapSort(input)
	}
}
