package GoSort

import (
	"testing"
)

func BenchmarkMergep1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := testDataSorted()
		MergeSortP(input)
	}
}

func BenchmarkMergep2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := testDataReversed()
		MergeSortP(input)
	}
}

func BenchmarkMergep3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := testDataRandom()
		MergeSortP(input)
	}
}
