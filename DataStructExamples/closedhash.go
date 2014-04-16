package main

import (
	"GoHash/ClosedHash"
	"fmt"
	// "math"
)

func main() {
	h := ClosedHash.NewHashTable(10)
	var p [5]*ClosedHash.ElemType
	ClosedHash.HashType = 1
	p[0] = ClosedHash.NewElem(62, 1)
	p[1] = ClosedHash.NewElem(17, 2)
	p[2] = ClosedHash.NewElem(39, 3)
	p[3] = ClosedHash.NewElem(29, 4)
	p[4] = ClosedHash.NewElem(9, 5)

	for i := 0; i < 5; i++ {
		h.Insert(*p[i])
	}
	fmt.Println(h)
	h.Traverse(ClosedHash.Print)
	h.Find(17)
	h.Delete(17)
	h.Traverse(ClosedHash.Print)

}
