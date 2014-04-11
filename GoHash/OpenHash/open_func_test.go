package OpenHash

import (
	"math/rand"
	"testing"
)

var size int = 100

func funcinti(factor int) *HashTable {
	ch := NewHashTable(size)
	for i := 0; i < factor; i++ {
		k := rand.Int()
		d := rand.Int()
		ce := NewElem(k, d)
		ch.Insert(*ce)
	}
	return ch
}

func BenchmarkHashDiv(b *testing.B) {
	ch := funcinti(30)
	HashType = 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch.Find(rand.Int())
	}
}

func BenchmarkHashMul(b *testing.B) {
	ch := funcinti(30)
	HashType = 1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch.Find(rand.Int())
	}
}
