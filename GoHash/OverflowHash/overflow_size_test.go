package OverflowHash

import (
	"math/rand"
	"testing"
)

var factor float32 = 0.3

func sizeinit(size int) *HashTable {
	ch := NewHashTable(size)
	for i := 0; i < int(float32(size)*factor); i++ {
		k := rand.Int() / size
		d := rand.Int()
		ce := NewElem(k, d)
		ch.Insert(*ce)
	}
	return ch
}
func BenchmarkSize100(b *testing.B) {
	// b.SetBytes(1024)
	ch := sizeinit(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch.Find(rand.Int() / factorsize)
	}
}

func BenchmarkSize1000(b *testing.B) {
	ch := sizeinit(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch.Find(rand.Int() / factorsize)
	}
}

func BenchmarkSize10000(b *testing.B) {
	ch := sizeinit(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch.Find(rand.Int() / factorsize)
	}
}

func BenchmarkSize100000(b *testing.B) {
	ch := sizeinit(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch.Find(rand.Int() / factorsize)
	}
}

func BenchmarkSize1000000(b *testing.B) {
	ch := sizeinit(1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch.Find(rand.Int() / factorsize)
	}
}
