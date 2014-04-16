package GoSort

import (
	"testing"
)

func BenchmarkBucket1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := testDataSorted()
		BucketSort(input)
	}
}

func BenchmarkBucket2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := testDataReversed()
		BucketSort(input)
	}
}

func BenchmarkBucket3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := testDataRandom()
		BucketSort(input)
	}
}
