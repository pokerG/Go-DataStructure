//This package provides some basically operate of linearlist which use array struct
//
//Copytright (C) 2014 by pokerG <pokerfacehlg@gmail.com>
package ArrayList

import (
	"errors"
	"fmt"
	"reflect"
)

type Pos interface{}
type Datatype interface{}

//struct's defaultsize
const DefaultSize = 20

//struct's actual sizeconst defaultSize = 20
var Size int

//Use array to achieve linearlist
type ArrayList struct {
	data    []Datatype
	lastPos int
}

//NewArrayListSize creates and initializes a new arraylist which have the size you give.
func NewArrayListSize(size int) *ArrayList {
	return &ArrayList{make([]Datatype, size), -1}
}

//NewArrayList creates and initializes a new arraylist which have the default size
func NewArrayList() *ArrayList {
	return NewArrayListSize(DefaultSize)
}

//MakeFull let the list is full
func (this *ArrayList) MakeFull() {
	this.lastPos = -1
}

//Insert have two arguements
// x is the value of element you want insert  p is the postion
//it may return three kind of error
//only return nil is successful
func (this *ArrayList) Insert(x Datatype, p Pos) error {
	v := reflect.ValueOf(p)
	v1 := int(v.Int())
	if v1 >= len(this.data)-1 {
		return errors.New("Over the index!")
	} else if v1 > this.lastPos+1 || v1 < 0 {
		return errors.New("This postion is not exist!")
	} else {
		for i := this.lastPos; i >= v1; i-- {
			this.data[i+1] = this.data[i]
		}
		this.data[v1] = x
		this.lastPos++
		return nil
	}
}

//Delete can delte the element of the postion you give
//it may return two kind of error
//only return nil is successful
func (this *ArrayList) Delete(p Pos) error {
	v := reflect.ValueOf(p)
	v1 := int(v.Int())
	if v1 > this.lastPos || v1 < 0 {
		return errors.New("The postion is not exist!")
	} else {
		this.lastPos--
		for i := v1; i <= this.lastPos; i++ {
			this.data[i] = this.data[i+1]
		}
		return nil
	}

}

//Find have one arguement
//x is the element's value of you want to find
//it return the postion and error
//postion = - 1 is not found
func (this *ArrayList) Find(x Datatype) (Pos, error) {
	for i := 0; i <= this.lastPos; i++ {
		if x == this.data[i] {
			return i, nil
		}
	}
	return -1, errors.New("Not found this element!")
}

//Print can print the elements of list
func (this *ArrayList) Print() {
	fmt.Println("Elements of List: ")
	for i := 0; i <= this.lastPos; i++ {
		fmt.Println(this.data[i], " ")
	}
	fmt.Println("")

}
