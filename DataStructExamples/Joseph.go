package main

import (
	// "errors"
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkList struct {
	phead *Node
	ptail *Node
	len   int
}

func GetNode(i int) *Node {
	pNode := &Node{i, nil}
	// if pNode != nil {
	// 	err := errors.New("Error,the memory is not enough!")
	// 	panic(err)
	// }
	return pNode
}

func (this *LinkList) Init() {
	p := GetNode(1)
	this.phead = p
	this.ptail = p
	p.next = this.phead
	this.len = 1
}

func (this *LinkList) Insert(n int) {

	i := 0
	var pNew *Node
	for i = 2; i <= n; i++ {
		pNew = GetNode(i)
		this.ptail.next = pNew
		this.ptail = pNew
		pNew.next = this.phead
		this.len++
	}
}

func (this *LinkList) Print() {
	pCur := this.phead
	for {
		fmt.Printf("The %d person.\n", pCur.data)
		pCur = pCur.next
		if pCur == this.phead {
			break
		}

	}
	fmt.Printf("The length of the List: %d\n", this.len)
}

func joseph(plist *LinkList, m int) {
	pPre := plist.ptail
	pCur := plist.phead
	var i int
	if plist.len != 1 {
		i = 0
	}

	for i < m-1 {
		pPre = pPre.next
		i++
	}
	pCur = pPre.next
	pPre.next = pCur.next
	plist.len--
	fmt.Printf("The last one is: %d\n", pPre.data)
}

func main() {
	n := 0
	pList := &LinkList{}
	fmt.Println("Please input the length of the Circle list: ")
	fmt.Scanf("%d\n", &n)
	m := 0
	fmt.Println("Please input the stop point: ")
	fmt.Scanf("%d\n", &m)
	pList.Init()
	pList.Insert(n)
	pList.Print()
	joseph(pList, m)
}
