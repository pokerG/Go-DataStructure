//This package provides some basically operete of stack
//
//Copyright (C) 2014 by pokerG <pokerfacehlg@gmail.com>
package Stack

import (
	"errors"
	"fmt"
)

//stack's defaultsize
const defaultSize = 20

//the stack struct
type GoStack struct {
	s       []interface{}
	top     int
	maxSize int
}

//NewGoStackSize creates and initializes a new GoStack
//which have the size you give
func NewGoStackSize(size int) *GoStack {
	return &GoStack{make([]interface{}, size), size, size}
}

//NewGoStack creates and initializes a new GoStack
//which have the default size
func NewGoStack() *GoStack {
	return NewGoStackSize(defaultSize)
}

//MakeFull let the stack is full
func (a *GoStack) MakeNull() {
	a.top = a.maxSize
}

//Empty if the stack is not empty return nil
func (a GoStack) Empty() error {
	if a.top > a.maxSize-1 {
		return errors.New("The Stack is empty")
	} else {
		return nil
	}
}

//Top return the element at the stack top
//but not delete it
func (a GoStack) Top() (interface{}, error) {
	if a.Empty() != nil {
		return 0, a.Empty()
	}
	return a.s[a.top], nil
}

//Pop return the element at the stact top
//and delete it
func (a *GoStack) Pop() (interface{}, error) {
	if a.Empty() != nil {
		return 0, a.Empty()
	}
	a.top = a.top + 1
	return a.s[a.top-1], nil
}

//Push make a new element into the stack
func (a *GoStack) Push(x interface{}) error {
	if a.top == 0 {
		return errors.New("The Stack is full!")
	} else {
		a.top = a.top - 1
		a.s[a.top] = x
		return nil
	}
}

//Print from the floor to ceil
func (a GoStack) Print() {
	i := a.top
	for {
		if i == a.maxSize {
			break
		}

		fmt.Print(a.s[i])
		i = i + 1

	}
}
