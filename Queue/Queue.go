package Queue

import (
	"errors"
	"fmt"
)

const defaultSize = 20

var qSize int

type GoQueue struct {
	q       []interface{}
	front   int
	rear    int
	maxSize int
}

func NewGoQueueSize(size int) *GoQueue {
	qSize = size
	return &GoQueue{make([]interface{}, size), 0, size - 1, size}
}

func NewGoQueue() *GoQueue {
	return NewGoQueueSize(defaultSize)
}

func addone(i int) int {
	return ((i + 1) % qSize)
}

func (a *GoQueue) MakeNull() {
	a.front = 0
	a.rear = a.maxSize - 1
}

func (a GoQueue) Empty() error {
	if addone(a.rear) == a.front {
		return errors.New("The Queue is empty!")
	} else {
		return nil
	}
}

func (a GoQueue) Front() (interface{}, error) {
	if a.Empty() != nil {
		return 0, a.Empty()
	}
	return a.q[a.front], nil
}

func (a *GoQueue) Enter(x interface{}) error {
	if addone(addone(a.rear)) == a.front {
		return errors.New("The Queue is full!")
	}
	a.rear = addone(a.rear)
	a.q[a.rear] = x
	return nil
}

func (a *GoQueue) Delete() (interface{}, error) {
	if a.Empty() != nil {
		return 0, a.Empty()
	}
	i := a.front
	a.front = addone(a.front)
	return a.q[i], nil
}

func (a GoQueue) Print() {
	for i := 0; i < qSize; i++ {
		fmt.Print(a.Front() + " ")
		x, _ := a.Delete()
		a.Enter(x)
	}

}
