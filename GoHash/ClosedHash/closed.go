//This package achieve the closed hash table
//
//Copytright (C) 2014 by pokerG <pokerfacehlg@gmail.com>
package ClosedHash

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
	key int
	ord int
}

type HashTable struct {
	elem      []ElemType
	count     int // current number of data
	sizeindex int // current size
}

func hash(k int) int {
	// fmt.Println(HashType)
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

func collision(p *int, d int) {
	(*p) = (*p + d) % HashSize
}

func (this *HashTable) search(k int, c *int) (int, error) {
	p := hash(k)
	// fmt.Println(k, p)
	for this.elem[p].key != 0 && !(k == this.elem[p].key) {
		(*c)++
		// fmt.Println(*c)
		if *c < HashSize {
			collision(&p, *c) //deal with collision
		} else {
			break
		}
	}

	if k == this.elem[p].key {
		return p, nil
	} else {
		return p, Unfind
	}
}

//Insert insert a new element into hash table
func (this *HashTable) Insert(e ElemType) error {
	c := 0
	p, err := this.search(e.key, &c)
	if err == nil {
		return Duplicate
	} else {
		this.elem[p] = e
		this.count++
		// fmt.Println(*this)
		return nil

	}

}

//Find search k from hash table
func (this *HashTable) Find(k int) bool {
	c := 0
	// p, err := this.search(k, &c)
	_, err := this.search(k, &c)
	if err == nil {
		// fmt.Printf("The element is %v  , address is %d\n", this.elem[p], p)
		return true
	} else {
		// fmt.Println(err)
		return false
	}
}

//Delete delete k from hash table
func (this *HashTable) Delete(k int) error {
	c := 0
	p, err := this.search(k, &c)
	if err == nil {
		this.elem[p].key = 0
		this.elem[p].ord = 0
		return nil
	} else {
		return err
	}
}

//Traverse print the hash table
func (this *HashTable) Traverse(Vi func(int, ElemType)) {
	fmt.Printf("Hase addr 0 ~ %d\n", HashSize)
	for i := 0; i < HashSize; i++ {
		if this.elem[i].key != 0 {
			Vi(i, this.elem[i])
		}
	}
}

func Print(p int, r ElemType) {
	fmt.Printf("address = %d (%d,%d)\n", p, r.key, r.ord)
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
	return &ElemType{key, ord}
}
