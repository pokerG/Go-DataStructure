package LinkedList

import (
	"fmt"
)

type Datatype interface{}

type Node struct {
	Data Datatype
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func NewNode(Data Datatype, Next *Node) *Node {
	return &Node{Data, Next}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

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
