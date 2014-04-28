package GoSort

import (
	"testing"
)

func BenchmarkQuickp1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := testDataSorted()
		QuickSortP(input)

	}
}

func BenchmarkQuickp2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := testDataReversed()
		QuickSortP(input)

	}
}

func BenchmarkQuickp3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := testDataRandom()
		QuickSortP(input)
	}
}
