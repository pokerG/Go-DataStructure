package main

import (
	"GoHash/OverflowHash"
	"fmt"
	// "math"
)

func main() {
	h := OverflowHash.NewHashTable(10)
	var p [5]*OverflowHash.ElemType
	p[0] = OverflowHash.NewElem(62, 1)
	p[1] = OverflowHash.NewElem(17, 2)
	p[2] = OverflowHash.NewElem(39, 3)
	p[3] = OverflowHash.NewElem(29, 4)
	p[4] = OverflowHash.NewElem(9, 5)

	for i := 0; i < 5; i++ {
		h.Insert(*p[i])
	}
	fmt.Println(h)
	h.Traverse(OverflowHash.Print)
	h.Find(17)
	h.Delete(17)
	h.Find(17)
	h.Find(29)
	h.Find(9)
	h.Find(5)
	h.Traverse(OverflowHash.Print)

}
