package ClosedHash

import (
	"math/rand"
	"testing"
)

var factorsize int = 100000

func factorinti(factor int) *HashTable {
	ch := NewHashTable(factorsize)
	HashType = 0
	for i := 0; i < factor; i++ {
		k := rand.Int()
		d := rand.Int()
		ce := NewElem(k, d)
		ch.Insert(*ce)
	}
	return ch
}
func BenchmarkFactor10(b *testing.B) {
	// b.SetBytes(1024)
	ch := factorinti(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch.Find(rand.Int())
	}
}

func BenchmarkFactor30(b *testing.B) {
	ch := factorinti(30000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch.Find(rand.Int())
	}
}

func BenchmarkFactor50(b *testing.B) {
	ch := factorinti(50000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch.Find(rand.Int())
	}
}

func BenchmarkFactor70(b *testing.B) {
	ch := factorinti(70000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch.Find(rand.Int())
	}
}

func BenchmarkFactor90(b *testing.B) {
	ch := factorinti(90000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch.Find(rand.Int() / factorsize)
	}
}
