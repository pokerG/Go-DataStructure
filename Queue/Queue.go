//This package provides some basically operete of queue
//
//Copyright (C) 2014 by pokerG <pokerfacehlg@gmail.com>
package Queue

import (
	"errors"
	"fmt"
)

const defaultSize = 20

var qSize int

//the queue struct
type GoQueue struct {
	q       []interface{}
	front   int
	rear    int
	maxSize int
}

//NewGoQueueSize creates and initializes a new GoQueue
//which have the size you give
func NewGoQueueSize(size int) *GoQueue {
	qSize = size
	return &GoQueue{make([]interface{}, size), 0, size - 1, size}
}

//NewGoQueue creates and initializes a new GoQueue
//which have the default size
func NewGoQueue() *GoQueue {
	return NewGoQueueSize(defaultSize)
}

func addone(i int) int {
	return ((i + 1) % qSize)
}

//MakeFull let the queue is full
func (a *GoQueue) MakeNull() {
	a.front = 0
	a.rear = a.maxSize - 1
}

//Empty if the queue is not empty return nil
func (a GoQueue) Empty() error {
	if addone(a.rear) == a.front {
		return errors.New("The Queue is empty!")
	} else {
		return nil
	}
}

//Front return the fisrt element of queue
//but not delete it
func (a GoQueue) Front() (interface{}, error) {
	if a.Empty() != nil {
		return 0, a.Empty()
	}
	return a.q[a.front], nil
}

//Push make a new element into the queue
func (a *GoQueue) Enter(x interface{}) error {
	if addone(addone(a.rear)) == a.front {
		return errors.New("The Queue is full!")
	}
	a.rear = addone(a.rear)
	a.q[a.rear] = x
	return nil
}

//Front return the fisrt element of queue
//and delete it
func (a *GoQueue) Delete() (interface{}, error) {
	if a.Empty() != nil {
		return 0, a.Empty()
	}
	i := a.front
	a.front = addone(a.front)
	return a.q[i], nil
}

//Print from the front to rear
func (a GoQueue) Print() {
	for i := 0; i < qSize; i++ {
		fmt.Print(a.Front() + " ")
		x, _ := a.Delete()
		a.Enter(x)
	}

}
