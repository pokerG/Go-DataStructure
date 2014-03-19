package main

import (
	"fmt"
)

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

func LeverOrder(t *CSNode) {
	var front, rear int
	front, rear = 0, 0
	var q [MAXLENGTH]*CSNode
	if t == nil {
		return
	}
	rear++
	q[rear] = t
	fmt.Printf("%s ", t.data)
	for front != rear {
		front++
		t = q[front].fchild
		for t != nil {
			fmt.Printf("%s ", t.data)
			rear++
			q[rear] = t
			t = t.rsib
		}
	}
	fmt.Println("")
}

func main() {
	//test data A B D # E I # # F # # C G # H # # # #
	t := CreateCT()
	fmt.Println("PreOrder:")
	PreOrder(t)
	fmt.Println("")
	fmt.Println("PostOrder:")
	PostOrder(t)
	fmt.Println("")
	fmt.Println("LeverOrder:")
	LeverOrder(t)
	fmt.Println("")
}
