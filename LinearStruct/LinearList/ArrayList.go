package LinearList

import (
	"LinearStruct"
	"errors"
	"fmt"
)

//Use array to achieve linearlist
type ArrayList struct {
	lastPos int
	data    []LinearStruct.Datatype
}

func NewArrayListSize(size int) *ArrayList {
	return &ArrayList{make([]LinearStruct.Datatype, size), -1}
}

func NewArrayList() *ArrayList {
	return NewArrayListSize(LinearStruct.DefaultSize)
}

func (this *ArrayList) MakeFull() {
	this.lastPos = -1
}

func (this *ArrayList) Insert(x LinearStruct.Datatype, p LinearStruct.Pos) error {
	if p >= len(this.data)-1 {
		return errors.New("Over the index!")
	} else if p > this.lastPos+1 || p < 0 {
		return errors.New("This postion is not exist!")
	} else {
		for i = this.lastPos; i >= p; i-- {
			this.data[i+1] = this.data[i]
		}
		this.data[p] = x
		this.lastPos++
		return nil
	}
}

func (this *ArrayList) Delete(p LinearStruct.Pos) error {
	if p > this.lastPos || p < 0 {
		return errors.New("The postion is not exist!")
	} else {
		this.lastPos--
		for i = p; i <= this.lastPos; i++ {
			this.data[i] = this.data[i+1]
		}
		return nil
	}

}

func (this *ArrayList) Find(x LinearStruct.Datatype) {
	for i = 0; i <= this.lastPos; i++ {
		if x == this.data[i] {
			return i, nil
		}
	}
	return -1, errors.New("Not found this element!")
}

func (this *ArrayList) Print() {
	fmt.Print("Elements of List: ")
	for i = 0; i <= this.lastPos; i++ {
		fmt.Print(this.data[i], " ")
	}

}
