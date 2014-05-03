//This package provides some basically operate of linearlist which use linked list
//
//Copytright (C) 2014 by pokerG <pokerfacehlg@gmail.com>
package LinkedList

import (
	"fmt"
)

type Datatype interface{}

type Node struct {
	Data Datatype
	Next *Node
}

//Use linked list to achieve linearlist
type LinkedList struct {
	Head *Node
}

//NewNode creates a new node use in linkedlist
//it have two arguements
//Data is the value of element Next is the postion in linkedlist
func NewNode(Data Datatype, Next *Node) *Node {
	return &Node{Data, Next}
}

//NewLinkedList creates and initializes a new linkedlist
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

//Insert can insert a node to the linkedlist
func (this *LinkedList) Insert(n Node) {
	if (*this).Head == nil {
		(*this).Head = &n
		return
	}

	ptr := (*this).Head
	for (*ptr).Next != nil {
		ptr = (*ptr).Next
	}
	(*ptr).Next = &n
}

//Find can find the fist element which value = dat
//return the node and successful sign
//return  nil and false is not found
func (this *LinkedList) Find(dat Datatype) (*Node, bool) {
	ptr := (*this).Head

	for (*ptr).Next != nil {
		if (*ptr).Data == dat {
			return ptr, true
		}
		ptr = (*ptr).Next
	}
	if (*ptr).Data == dat {
		return ptr, true
	}
	return nil, false
}

//Delete can delete the first element which value = dat
//only if return nil  is delete successful
func (this *LinkedList) Delete(dat Datatype) error {
	_, b := this.Find(dat)
	if !b {
		err := fmt.Errorf("**ERROR** Data: %d not found in list", dat)
		return err
	}

	if (*this).Head.Data == dat {
		(*this).Head = (*this).Head.Next
	}

	ptr := (*this).Head
	for (*ptr).Next != nil {
		if (*ptr).Next.Data == dat && (*ptr).Next.Next != nil {
			(*ptr).Next = (*ptr).Next.Next
		} else if (*ptr).Next.Data == dat && (*ptr).Next.Next == nil {
			(*ptr).Next = nil
			return nil
		}
		ptr = (*ptr).Next
	}
	return nil
}

//Print can print the elements of list
func (this *LinkedList) Print() {
	if (*this).Head == nil {
		fmt.Println("**WARNING** Printing out an empty list")
		return
	}

	ptr := (*this).Head
	fmt.Println("Linked List: ")
	for (*ptr).Next != nil {
		fmt.Printf("Node %v -> ", (*ptr).Data)
		ptr = (*ptr).Next
	}
	fmt.Println("Node", (*ptr).Data)
	return
}
