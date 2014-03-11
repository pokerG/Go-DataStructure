// Copyright 2014 pokerface. All rights reserved.

/*
Package LinearStruct provides basic interfaces of linearstruct.
Such as linearlist , queue , stack , array.
*/
package LinearStruct

type Pos interface{}
type Datatype interface{}

//struct's defaultsize
const DefaultSize = 20

//struct's actual sizeconst defaultSize = 20
var Size int

//LinearList provides basic operations of linearlist
type LinearList interface {
	Insert(Datatype, Pos) error
	Delete(Pos) error
	MakeNull()
	Find(Datatype) (Pos, error)
	Print()
}

//Queue provides basic operations of queue
type Queue interface {
	MakeNull()
	Empty() error
	Front() (Datatype, error)
	Enter() (Datatype, error)
	Delete() (Datatype, error)
	Print()
}

//Stack provides basic operations of stack
type Stack interface {
	MakeNull()
	Empty() error
	Top() (Datatype, error)
	Push(Datatype) error
	Pop() (Datatype, error)
	Print()
}
