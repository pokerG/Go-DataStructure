package main

import (
	"GoHash/OpenHash"
	// "fmt"
	"math/rand"
)

var size int = 100000

func funcinti(factor int) *OpenHash.HashTable {
	ch := OpenHash.NewHashTable(size)
	for i := 0; i < factor; i++ {
		k := rand.Int()
		d := rand.Int()
		ce := OpenHash.NewElem(k, d)
		ch.Insert(*ce)
	}
	return ch
}

func main() {
	OpenHash.HashType = 2
	h := funcinti(3000)

	for i := 0; i < 3000; i++ {
		h.Find(rand.Int())
	}

}
