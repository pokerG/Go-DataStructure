package GoSort

import (
	"testing"
)

func BenchmarkInsert1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := testDataSorted()
		InsertSort(input)
	}
}

func BenchmarkInsert2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := testDataReversed()
		InsertSort(input)
	}
}

func BenchmarkInsert3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := testDataRandom()
		InsertSort(input)
	}
}
