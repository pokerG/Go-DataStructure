package GoSort

import (
	"testing"
)

func BenchmarkSelect1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := testDataSorted()
		SelectSort(input)
	}
}

func BenchmarkSelect2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := testDataReversed()
		SelectSort(input)
	}
}

func BenchmarkSelect3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := testDataRandom()
		SelectSort(input)
	}
}
