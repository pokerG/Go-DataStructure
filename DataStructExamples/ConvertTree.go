/*
	Convert between forest and binaryTree
	use adjacency list
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAXSIZE int = 20

type CTNode struct { //forest child node
	child int
	next  *CTNode
}

type CTBox struct { //forest header node
	data       int
	firstchild *CTNode
}

type CTree struct { //forest
	nodes [MAXSIZE]CTBox
	n     int //the factual number of nodes
}

type Bnode struct { //binaryTree node
	data           int
	lchild, rchild *Bnode
}

func CreateBT() *Bnode {
	var ch string
	t := &Bnode{}
	fmt.Scanf("%s", &ch)
	if ch == "#" {
		t = nil
	} else {
		t.data, _ = strconv.Atoi(ch)
		t.lchild = CreateBT()
		t.rchild = CreateBT()
	}
	return t
}

func CreateCT() *CTree {
	t := &CTree{}
	t.n = 0
	fmt.Println("Please input the node data:(0 is end)")
	for {
		var ch int
		fmt.Scanf("%d", &ch)
		if ch != 0 {
			t.nodes[t.n].data = ch
			t.nodes[t.n].firstchild = nil
			t.n++
		}
	}
	for i := 0; i < t.n; i++ {
		fmt.Printf("Please input the %dth node's childs:")
		reader := bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()
		command := string(data)
		children := strings.Fields(command)
		for _, v := range children {
			tmp := t.nodes[i].firstchild
			data, _ := strconv.Atoi(v)
			n := &CTNode{data, nil}
			t.nodes[i].firstchild = n
			n.next = tmp
		}
	}
	return t
}

var visited [MAXSIZE]bool

func ForestToBt(ct *CTree, bt *Bnode) {
	if ct.n == 0 {
		bt = nil
		return
	}
	for i := 0; i <= ct.n; i++ {
		if !visited[i] {
			PreOrder(ct.nodes[i])
		}
	}

}
