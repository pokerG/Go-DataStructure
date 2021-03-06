//This package achieve the  hash table have overflow list
//
//Copytright (C) 2014 by pokerG <pokerfacehlg@gmail.com>
package OverflowHash

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
	switch HashType { // choose the hash function you want use
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
	// fmt.Println(k, p)
	tmp := &this.elem[p]
	for tmp != nil {
		if tmp.key == k {
			return p, tmp
		} else {
			tmp = tmp.next
		}
	}
	return p, tmp
}

//Insert insert a new element into hash table
func (this *HashTable) Insert(e ElemType) error {
	p, r := this.search(e.key)
	if r != nil {
		return Duplicate
	} else {
		tmp := &ElemType{}
		tmp.key = e.key
		tmp.ord = e.ord
		tmp.next = nil
		if this.elem[p].key == 0 {
			this.elem[p].key = e.key
			this.elem[p].ord = e.ord
		} else {
			tmp.next = this.elem[p].next
			this.elem[p].next = tmp
		}
		this.count++
		// fmt.Println(*this)
		return nil
	}

}

//Find search k from hash table
func (this *HashTable) Find(k int) bool {
	// p, r := this.search(k)
	_, r := this.search(k)
	if r != nil {
		// fmt.Printf("The element is %v  , address is %d\n", r, p)\
		return true
	} else {
		// fmt.Println(Unfind)
		return false
	}
}

//Delete delete k from hash table
func (this *HashTable) Delete(k int) error {
	p := hash(k)
	tmp := &this.elem[p]
	ftmp := &this.elem[p]
	for tmp != nil {
		if tmp.key == k {
			if tmp.next == nil {
				tmp.key = 0
				tmp.ord = 0
			} else if ftmp == tmp {
				tmp.key = tmp.next.key
				tmp.ord = tmp.next.ord
				// t := tmp.next.next
				tmp.next = tmp.next.next
			} else {
				ftmp.next = tmp.next
			}
			return nil
		}
		ftmp = tmp
		tmp = tmp.next
	}
	return Unfind

}

//Traverse print the hash table
func (this *HashTable) Traverse(Vi func(ElemType)) {
	fmt.Printf("Hase addr 0 ~ %d\n", HashSize)
	for i := 0; i < HashSize; i++ {
		if this.elem[i].key != 0 {
			Vi(this.elem[i])
		}
	}
}

func Print(r ElemType) {
	tmp := &r
	for tmp != nil {
		fmt.Printf("address = %d (%d,%d)\n", tmp, tmp.key, tmp.ord)
		tmp = tmp.next
	}
}

//NewHashTable creates and initializes a new hash table which have the size you give.
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

//NewElem creates a new element hash table required
func NewElem(key int, ord int) *ElemType {
	return &ElemType{key, ord, nil}
}
