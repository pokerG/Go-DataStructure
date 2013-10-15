package Stack

import (
	"errors"
)

const defaultSize = 20

type GoStack struct {
	s       []interface{}
	top     int
	maxSize int
}

func NewGoStackSize(size int) *GoStack {
	return &GoStack{make([]interface{}, size), size, size}
}

func NewGoStack() *GoStack {
	return NewGoStackSize(defaultSize)
}

func (a *GoStack) MakeNull() {
	a.top = a.maxSize
}

func (a GoStack) Empty() error {
	if a.top > a.maxSize - 1 {
		return errors.New("The Stack is empty")
	} else {
		return nil
	}
}

func (a GoStack) Top() (interface{}, error) {
	if a.Empty() != nil {
		return 0, a.Empty()
	}
	return a.s[a.top], nil
}

func (a *GoStack) Pop() (interface{}, error) {
	if a.Empty() != nil {
		return 0, a.Empty()
	}
	a.top = a.top + 1
	return a.s[a.top - 1], nil
}

func (a *GoStack) Push(x interface{}) error {
	if a.top == 0 {
		return errors.New("The Stack is full!")
	} else {
		a.top = a.top - 1
		a.s[a.top] = x
		return nil
	}
}
