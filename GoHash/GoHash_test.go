package GoHash

import (
	"math/rand"
	"testing"
)

var m = 100000

func BenchmarkHashDiv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HashDiv(rand.Int(), m)
	}
}

func BenchmarkHashMul(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HashMul(rand.Int(), m)
	}
}
