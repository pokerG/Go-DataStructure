package GoSort

import (
	"testing"
)

func BenchmarkBubble1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := testDataSorted()
		BubbleSort(input)
	}
}

func BenchmarkBubble2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := testDataReversed()
		BubbleSort(input)
	}
}

func BenchmarkBubble3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := testDataRandom()
		BubbleSort(input)
	}
}
