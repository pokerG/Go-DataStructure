package CursorList

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

func NewSpaceSize(size int) *Space {
	return &Space{make([]Node, size), size, 0, 0, 0}
}

func NewSpace() *Space {
	return NewSpaceSize(DefaultSize)
}

func (this *Space) Init() {
	var j int
	for j = 0; j < this.size-1; j++ {
		this.space[j].next = Cursor(j + 1)
	}
	this.space[j].next = -1
	this.avail = 0
}

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

func (this *Space) freeNode(q Cursor) {
	this.space[q].next = this.space[this.avail].next
	this.space[this.avail].next = q
}

func (this *Space) Insert(x Datatype, p Position) {
	var q Position
	q = Position(this.getNode())
	this.space[q].data = x
	this.space[q].next = this.space[p].next
	this.space[p].next = Cursor(q)
}

func (this *Space) Delete(p Position) {
	var q Position
	if this.space[p].next != -1 {
		q = Position(this.space[p].next)
		this.space[p].next = this.space[q].next
		this.freeNode(Cursor(q))
	}
}
