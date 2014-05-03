//Forest travel
//
//Copytright (C) 2014 by pokerG <pokerfacehlg@gmail.com>
package main

import (
	"fmt"
)

//test data A B D # E I # # F # # C G # H # # # #

const MAXLENGTH int = 20

type CSNode struct {
	data         string
	fchild, rsib *CSNode
}

func CreateCT() *CSNode {
	var ch string
	fmt.Scanf("%s", &ch)
	nd := &CSNode{}

	if ch == "#" {
		nd = nil
	} else {
		nd.data = ch
		nd.fchild = CreateCT()
		nd.rsib = CreateCT()
	}
	// fmt.Println(ch, nd)
	return nd
}

func PreOrder(t *CSNode) {
	if t != nil {
		fmt.Printf("%s ", t.data)
		PreOrder(t.fchild)
		PreOrder(t.rsib)
	}
}

func PostOrder(t *CSNode) {
	if t != nil {
		PostOrder(t.fchild)
		PostOrder(t.rsib)
		fmt.Printf("%s ", t.data)
	}
}

func LeverOrder(t []*CSNode) {
	var front, rear int
	front, rear = 0, 0
	var q [MAXLENGTH]*CSNode
	if t[0] == nil {
		return
	}
	for _, v := range t {
		rear++
		q[rear] = v
		fmt.Printf("%s ", v.data)
	}
	var tmp *CSNode
	for front != rear {
		front++
		tmp = q[front].fchild
		for tmp != nil {
			fmt.Printf("%s ", tmp.data)
			rear++
			q[rear] = tmp
			tmp = tmp.rsib
		}
	}
	fmt.Println("")
}

func main() {
	var t []*CSNode
	var n int
	fmt.Println("Input the number of tree:")
	fmt.Scanf("%d\n", &n)
	t = make([]*CSNode, n)
	for i, _ := range t {
		t[i] = CreateCT()
	}
	fmt.Println("PreOrder:")
	for _, v := range t {
		PreOrder(v)
		fmt.Println("")
	}

	fmt.Println("PostOrder:")
	for _, v := range t {
		PostOrder(v)
		fmt.Println("")
	}
	fmt.Println("LeverOrder:")
	LeverOrder(t)
	fmt.Println("")
}
