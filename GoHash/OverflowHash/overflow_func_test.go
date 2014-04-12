package OverflowHash

import (
	"math/rand"
	"testing"
)

var size int = 100000

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
	HashType = 0
	ch := funcinti(30000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch.Find(rand.Int())
	}
}

func BenchmarkHashMul(b *testing.B) {
	HashType = 1
	ch := funcinti(30000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch.Find(rand.Int())
	}
}

func BenchmarkHashBit(b *testing.B) {
	HashType = 2
	ch := funcinti(30000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch.Find(rand.Int())
	}
}
