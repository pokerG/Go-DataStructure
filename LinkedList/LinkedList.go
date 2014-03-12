package LinkedList

import (
	"fmt"
)

type Datatype interface{}

type Node struct {
	data Datatype
	next *Node
}

type LinkedList struct {
	head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (this *LinkedList) Insert(n Node) {
	if (*this).head == nil {
		(*this).head = &n
		return
	}

	ptr := (*this).head
	for (*ptr).next != nil {
		ptr = (*ptr).next
	}
	(*ptr).next = &n
}

func (this *LinkedList) Find(dat Datatype) (*Node, bool) {
	ptr := (*this).head

	for (*ptr).next != nil {
		if (*ptr).data == dat {
			return ptr, true
		}
		ptr = (*ptr).next
	}
	if (*ptr).data == dat {
		return ptr, true
	}
	return nil, false
}

func (this *LinkedList) Delete(dat Datatype) error {
	_, b := this.Find(dat)
	if !b {
		err := fmt.Errorf("**ERROR** Data: %d not found in list", dat)
		return err
	}

	if (*this).head.data == dat {
		(*this).head = (*this).head.next
	}

	ptr := (*this).head
	for (*ptr).next != nil {
		if (*ptr).next.data == dat && (*ptr).next.next != nil {
			(*ptr).next = (*ptr).next.next
		} else if (*ptr).next.data == dat && (*ptr).next.next == nil {
			(*ptr).next = nil
			return nil
		}
		ptr = (*ptr).next
	}
	return nil
}

func (this *LinkedList) Print() {
	if (*this).head == nil {
		fmt.Println("**WARNING** Printing out an empty list")
		return
	}

	ptr := (*this).head
	fmt.Println("Linked List: ")
	for (*ptr).next != nil {
		fmt.Printf("Node %d -> ", (*ptr).data)
		ptr = (*ptr).next
	}
	fmt.Println("Node", (*ptr).data)
	return
}
