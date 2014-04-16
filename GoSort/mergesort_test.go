package GoSort

import (
	"testing"
)

func BenchmarkMerge1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := testDataSorted()
		MergeSort(input)
	}
}

func BenchmarkMerge2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := testDataReversed()
		MergeSort(input)
	}
}

func BenchmarkMerge3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := testDataRandom()
		MergeSort(input)
	}
}
