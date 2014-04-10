package OverflowHash

import (
	"math/rand"
	"testing"
)

var factorsize int = 100

func factorinti(factor int) *HashTable {
	ch := NewHashTable(factorsize)
	for i := 0; i < factor; i++ {
		k := rand.Int() / factorsize
		d := rand.Int()
		ce := NewElem(k, d)
		ch.Insert(*ce)
	}
	return ch
}
func BenchmarkFactor10(b *testing.B) {
	// b.SetBytes(1024)
	ch := factorinti(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch.Find(rand.Int() / factorsize)
	}
}

func BenchmarkFactor30(b *testing.B) {
	ch := factorinti(30)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch.Find(rand.Int() / factorsize)
	}
}

func BenchmarkFactor50(b *testing.B) {
	ch := factorinti(50)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch.Find(rand.Int() / factorsize)
	}
}

func BenchmarkFactor70(b *testing.B) {
	ch := factorinti(70)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch.Find(rand.Int() / factorsize)
	}
}

func BenchmarkFactor90(b *testing.B) {
	ch := factorinti(90)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch.Find(rand.Int() / factorsize)
	}
}
