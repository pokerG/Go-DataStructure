//This package provides some basically operate of linearlist which use cursor
//
//Copytright (C) 2014 by pokerG <pokerfacehlg@gmail.com>
package CursorList

import "fmt"

type Datatype interface{}
type Cursor int
type Position int

const DefaultSize = 20

type Node struct {
	data Datatype
	next Cursor
}

type Space struct {
	space []Node
	size  int
	avail Cursor
	l     Cursor
	m     Cursor
}

//NewSpacesize creates and initiallizes a new cursorlist which have the size you give
func NewSpaceSize(size int) *Space {
	return &Space{make([]Node, size), size, 0, 0, 0}
}

//NewSpace creates and initializes a new cursorlist which have the defualt size
func NewSpace() *Space {
	return NewSpaceSize(DefaultSize)
}

//Init the cursorlist
func (this *Space) Init() {
	var j int
	for j = 0; j < this.size-1; j++ {
		this.space[j].next = Cursor(j + 1)
	}
	this.space[j].next = -1
	this.avail = 0
}

//getNode make a new area in space
func (this *Space) getNode() Cursor {
	var p Cursor
	if this.space[this.avail].next == -1 {
		p = -1
	} else {
		p = this.space[this.avail].next
		this.space[this.avail].next = this.space[p].next
	}
	return p
}

//freeNode delete the area in space
func (this *Space) freeNode(q Cursor) {
	this.space[q].next = this.space[this.avail].next
	this.space[this.avail].next = q
}

//Insert have two arguements
// x is the value of element you want insert  p is the postion
func (this *Space) Insert(x Datatype, p Position) {
	var q Position
	q = Position(this.getNode())
	this.space[q].data = x
	this.space[q].next = this.space[p].next
	this.space[p].next = Cursor(q)
}

//Delete can delte the element of the postion you give
func (this *Space) Delete(p Position) {
	var q Position
	if this.space[p].next != -1 {
		q = Position(this.space[p].next)
		this.space[p].next = this.space[q].next
		this.freeNode(Cursor(q))
	}
}

//Print can print the elements of cursorlist
func (this *Space) Print() {
	for i := 0; i < this.size; i++ {
		fmt.Printf("%d--->%d ## %d\n", i, this.space[i].data, this.space[i].next)
	}
}
