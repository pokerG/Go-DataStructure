package GoSort

import (
	"testing"
)

func BenchmarkShell1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := testDataSorted()
		ShellSort(input)
	}
}

func BenchmarkShell2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := testDataReversed()
		ShellSort(input)
	}
}

func BenchmarkShell3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := testDataRandom()
		ShellSort(input)
	}
}
