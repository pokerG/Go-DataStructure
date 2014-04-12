package OpenHash

import (
	"GoHash"
	"errors"
	"fmt"
)

const N int = 100 // the size of data
var HashSize int  // the size of hash table
var HashType int  // the hash func
var (
	Duplicate = errors.New("There already have the element!")
	Unfind    = errors.New("Not Found!")
)

type ElemType struct {
	key  int
	ord  int
	next *ElemType
}

type HashTable struct {
	elem      []ElemType
	count     int // current number of data
	sizeindex int // current size
}

func hash(k int) int {
	switch HashType {
	case 0:
		return GoHash.HashDiv(k, HashSize)
	case 1:
		return GoHash.HashMul(k, HashSize)
	case 2:
		return GoHash.HashBit(k, HashSize)
	default:
		return GoHash.HashDiv(k, HashSize)
	}

}

func (this *HashTable) search(k int) (int, *ElemType) {
	p := hash(k)
	fmt.Println(k, p)
	tmp := this.elem[p].next
	for tmp != nil {
		if tmp.key == k {
			return p, tmp
		} else {
			tmp = tmp.next
		}
	}
	return p, tmp
}

func (this *HashTable) Insert(e ElemType) error {
	p, r := this.search(e.key)
	if r != nil {
		return Duplicate
	} else {
		tmp := &ElemType{}
		tmp.key = e.key
		tmp.ord = e.ord
		tmp.next = nil
		tmp.next = this.elem[p].next
		this.elem[p].next = tmp
		this.count++
		// fmt.Println(*this)
		return nil
	}

}

func (this *HashTable) Find(k int) bool {
	// p, r := this.search(k)
	_, r := this.search(k)
	if r != nil {
		// fmt.Printf("The element is %v  , address is %d\n", r, p)
		return true
	} else {
		// fmt.Println(Unfind)
		return false
	}
}

func (this *HashTable) Delete(k int) error {
	p := hash(k)
	tmp := this.elem[p].next
	ftmp := &this.elem[p]
	for tmp != nil {
		if tmp.key == k {
			ftmp.next = tmp.next
			return nil
		}
		ftmp = tmp
		tmp = tmp.next
	}
	return Unfind

}

func (this *HashTable) Traverse(Vi func(ElemType)) {
	fmt.Printf("Hase addr 0 ~ %d\n", HashSize)
	for i := 0; i < HashSize; i++ {
		if this.elem[i].next != nil {
			Vi(this.elem[i])
		}
	}
}

func Print(r ElemType) {
	tmp := r.next
	for tmp != nil {
		fmt.Printf("address = %d (%d,%d)\n", tmp, tmp.key, tmp.ord)
		tmp = tmp.next
	}
}

func NewHashTable(size int) *HashTable {
	h := &HashTable{}
	h.count = 0
	h.sizeindex = 0
	HashSize = size
	h.elem = make([]ElemType, HashSize)
	for i := 0; i < HashSize; i++ {
		h.elem[i].key = 0
	}
	return h
}

func NewElem(key int, ord int) *ElemType {
	return &ElemType{key, ord, nil}
}
