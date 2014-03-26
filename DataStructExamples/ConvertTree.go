/*
	Convert between forest and binaryTree
	use adjacency list
*/

package main

import (
	// "bufio"
	"fmt"
	// "os"
	// "strconv"
	// "strings"
)

const MAXSIZE int = 20

type CTNode struct { //forest child node
	sign int
	next *CTNode
}

type CTBox struct { //forest header node
	data       string
	parent     int
	firstchild *CTNode
}

type CTree struct { //forest
	nodes [MAXSIZE]CTBox
	n     int //the factual number of nodes
}

type Bnode struct { //binaryTree node
	data string
	// sign           int
	lchild, rchild *Bnode
}

var visited [MAXSIZE]bool

/*
 test data:
a -1
b 0
c 0
k 2
d -1
e 4
f 4
# 0

a -1
b 0
c 0
d 1
e 1
f 1
g 2
h 2
i 4
# 0

A B D H # # I # # E # # C F # J # # G # #
*/

func main() {
	ct := CreateCT()

	bt := ForestToBt(ct, -1)
	BTPrint(bt)
	fmt.Println("")
	// ct := &CTree{}
	// bt := CreateBT()
	// BTPrint(bt)
	// fmt.Println("")
	// BtToForest(bt, ct, "")
	// for i := 0; i < ct.n; i++ {
	// 	for j := 0; j < ct.n; j++ {
	// 		if ct.nodes[j].parent == i {
	// 			tmp := ct.nodes[i].firstchild
	// 			n := &CTNode{j, nil}
	// 			ct.nodes[i].firstchild = n
	// 			n.next = tmp
	// 		}
	// 	}
	// }
	// CTPrint(ct)
}

func BtToForest(bt *Bnode, ct *CTree, s string) {
	if bt == nil {
		return
	}
	// fmt.Println("!!", bt.data)
	a := Find(ct, s)
	if a == -1 {
		ct.nodes[ct.n].data = bt.data
		ct.nodes[ct.n].parent = Find(ct, s)
		ct.n++
	}

	tmp := bt.lchild
	if tmp != nil {
		ct.nodes[ct.n].data = tmp.data
		ct.nodes[ct.n].parent = Find(ct, bt.data)
		ct.n++
		tmp = tmp.rchild
	}

	for tmp != nil {
		// fmt.Println(tmp.data)
		x := Find(ct, tmp.data)
		if x == -1 {
			ct.nodes[ct.n].data = tmp.data
			ct.nodes[ct.n].parent = Find(ct, bt.data)
			ct.n++
		}
		tmp = tmp.rchild
	}
	// fmt.Println(ct)
	BtToForest(bt.lchild, ct, bt.data)
	if s == "" {
		BtToForest(bt.rchild, ct, "")
	}
}

func Find(ct *CTree, s string) int {
	for i := 0; i < ct.n; i++ {
		if s == ct.nodes[i].data {
			return i
		}
	}
	return -1
}
func ForestToBt(ct *CTree, parent int) *Bnode {
	var tmp int = -1
	bt := &Bnode{}
	for i := 0; i < ct.n; i++ {
		if ct.nodes[i].parent == parent && !visited[i] {
			tmp = i
			break
		}
	}

	if tmp == -1 {
		return nil
	}
	visited[tmp] = true
	fmt.Println(tmp, ct.nodes[tmp].data)
	bt.data = ct.nodes[tmp].data

	bt.lchild = ForestToBt(ct, tmp)
	bt.rchild = ForestToBt(ct, parent)
	return bt
}

func CreateBT() *Bnode {
	var ch string
	t := &Bnode{}
	fmt.Scanf("%s", &ch)
	if ch == "#" {
		t = nil
	} else {
		t.data = ch
		t.lchild = CreateBT()
		t.rchild = CreateBT()
	}
	return t
}

func CreateCT() *CTree {
	t := &CTree{}
	t.n = 0
	fmt.Println("Please input the node data and its parent number(# 0 is end)")
	for {
		var ch1 string
		var ch2 int
		fmt.Scanf("%s %d\n", &ch1, &ch2)
		// fmt.Println(ch1, ch2)
		if ch1 != "#" || ch2 != 0 {
			t.nodes[t.n].data = ch1
			t.nodes[t.n].parent = ch2
			t.nodes[t.n].firstchild = nil
			t.n++
		} else {
			break
		}
	}
	// fmt.Println("!!!")
	for i := 0; i < t.n; i++ {
		for j := 0; j < t.n; j++ {
			if t.nodes[j].parent == i {
				tmp := t.nodes[i].firstchild
				n := &CTNode{j, nil}
				t.nodes[i].firstchild = n
				n.next = tmp
			}
		}
	}
	return t
}

func BTPrint(bt *Bnode) {
	if bt != nil {
		fmt.Print(bt.data)
		if bt.lchild != nil || bt.rchild != nil {
			fmt.Print("(")
			BTPrint(bt.lchild)
			if bt.rchild != nil {
				fmt.Print(",")
			}
			BTPrint(bt.rchild)
			fmt.Print(")")
		}
	}
}

func CTPrint(ct *CTree) {
	for i := 0; i < ct.n; i++ {
		fmt.Print(ct.nodes[i].data, " ", ct.nodes[i].parent, ":  ")
		tmp := ct.nodes[i].firstchild
		for tmp != nil {
			fmt.Print(ct.nodes[tmp.sign].data, " ")
			tmp = tmp.next
		}
		fmt.Println("")
	}
}
