package main

import (
	"GoHash/ClosedHash"
	"GoHash/OpenHash"
	"GoHash/OverflowHash"
	"fmt"
	"math/rand"
	"time"
	// "math"
)

func main() {
	fmt.Println("The fill factor test:")
	size := 100
	factor := 10n 

	for ; factor <= size; factor = factor + 10 {
		ch := ClosedHash.NewHashTable(size)
		oh := OpenHash.NewHashTable(size)
		ofh := OverflowHash.NewHashTable(size)
		rand.Seed(time.Now().Unix())
		for i := 0; i < factor; i++ {
			k := rand.Int() / 100
			d := rand.Int()
			ce := ClosedHash.NewElem(k, d)
			oe := OpenHash.NewElem(k, d)
			ofe := OverflowHash.NewElem(k, d)
			ch.Insert(ce)
			oh.Insert(oe)
			ofh.Insert(ofe)
		}
		ft := time.Now()
		for i := 0; i < 10000; i++{
			ch.Find(rand.Int()/100)
		}
	}
}
