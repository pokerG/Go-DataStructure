package GoSort

import (
	"testing"
)

func BenchmarkQuick1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := testDataSorted()
		QuickSort(input)
	}
}

func BenchmarkQuick2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := testDataReversed()
		QuickSort(input)
	}
}

func BenchmarkQuick3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := testDataRandom()
		QuickSort(input)
	}
}
