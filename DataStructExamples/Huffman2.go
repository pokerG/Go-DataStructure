//This is huffmancode example
//Use max heap
//Copytright (C) 2014 by pokerG <pokerfacehlg@gmail.com>
package main

import (
	"container/heap"
	"fmt"
)

/*
test data
5
a 0.15
b 0.45
c 0.25
d 0.20
e 0.08

6
a 0.15
b 0.45
c 0.25
d 0.20
e 0.08
f 0.80
*/

var n int = 5         //the number of leaf
var m int = (2*n - 1) //the number of node
const maxVal float64 = 10000.0
const maxSize int = 100 //huffman code's wightest digit

type HuffmanNode struct {
	ch                     string
	weight                 float64
	lchild, rchild, parent int
	local                  int //the poistion in array
	index                  int // the index in heap
}

//****** Down is priorityQueue ******
type priorityQueue []*HuffmanNode

func (this priorityQueue) Len() int {
	return len(this)
}

func (this priorityQueue) Less(i, j int) bool {
	return this[i].weight < this[j].weight //pop the lowest
}

func (this priorityQueue) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
	this[i].index = i
	this[j].index = j
}

func (this *priorityQueue) Push(x interface{}) {
	n := len(*this)
	node := x.(*HuffmanNode)
	node.index = n

	*this = append(*this, node)
}

func (this *priorityQueue) Pop() interface{} {
	old := *this
	n := len(old)
	node := old[n-1]
	node.index = -1 //for safety
	*this = old[0 : n-1]
	return node
}

func (this *priorityQueue) update(node *HuffmanNode, ch string, weight float64) {
	heap.Remove(this, node.index)
	node.ch = ch
	node.weight = weight
	heap.Push(this, node)

}

//****** Up is priorityQueue ******

type CodeType struct {
	bits  []byte //bit string
	start int    //code's start position in bits
	ch    string
}

//Huffman build huffman tree
func Huffman(tree []HuffmanNode) {
	var i int
	var f float64
	var c string
	for i = 0; i < m; i++ {
		tree[i].parent = 0
		tree[i].local = i
		tree[i].lchild = -1
		tree[i].rchild = -1
		tree[i].weight = 0.0
	}
	h := &priorityQueue{}
	heap.Init(h)

	fmt.Printf("Please input %d node's byte and weight(sepatate by space)\n", n)
	for i = 0; i < n; i++ {
		fmt.Printf("Input the %dth byte and weight: ", i+1)
		fmt.Scanf("%v %v\n", &c, &f)
		tree[i].ch = c
		tree[i].weight = f
		heap.Push(h, &tree[i])
	}
	for i = n; i < m; i++ {
		// for h.Len() > 0 {

		//Use max heap accelerate the process of build tree
		n1 := heap.Pop(h).(*HuffmanNode)
		n2 := heap.Pop(h).(*HuffmanNode)
		fmt.Println(n1, n2)
		tree[n1.local].parent = i
		tree[n2.local].parent = i
		tree[i].lchild = n1.local
		tree[i].rchild = n2.local
		tree[i].weight = tree[n1.local].weight + tree[n2.local].weight
		heap.Push(h, &tree[i])
	}
}

//HuffmanCode use huffman tree make huffman code
func HuffmanCode(code []CodeType, tree []HuffmanNode) {
	var i, c, p int

	for i = 0; i < n; i++ {
		var cd CodeType //buffer
		cd.bits = make([]byte, n)
		cd.start = n
		cd.ch = tree[i].ch
		c = i
		p = tree[i].parent
		for p != 0 {
			cd.start--
			if tree[p].lchild == c {
				cd.bits[cd.start] = '0'
			} else {
				cd.bits[cd.start] = '1'
			}
			c = p
			p = tree[p].parent
		}
		code[i] = cd
	}

}

//HuffmanDecode use huffman tree decode huffman code
func HuffmanDecode(tree []HuffmanNode) {
	var i, j int = 0, 0
	var s string
	i = m - 1
	fmt.Print("Input the code to send('2' is the end flag):")
	fmt.Scanf("%s\n", &s)
	b := []byte(s)
	for b[j] != '2' {
		if b[j] == '0' {
			i = tree[i].lchild
		} else {
			i = tree[i].rchild
		}
		if tree[i].lchild == -1 {
			fmt.Printf("%v", tree[i].ch)
			i = m - 1
		}
		j++
	}
	fmt.Println("")
	if tree[i].lchild != -1 && b[j] != '2' {
		fmt.Print("\nError\n")
	}
}

func main() {
	fmt.Println("          --Huffman code--")
	fmt.Print("Please input the number of characters :")
	fmt.Scan(&n)
	m = (2*n - 1)
	fmt.Printf("A total of %d characters\n", n)
	tree := make([]HuffmanNode, m)
	code := make([]CodeType, n)
	for i := 0; i < n; i++ {
		code[i].bits = make([]byte, n)
	}
	var i, j int
	Huffman(tree)
	HuffmanCode(code, tree)
	fmt.Println("Output character's huffman code")
	for i = 0; i < n; i++ {
		fmt.Printf("%v: ", code[i].ch)
		for j = code[i].start; j < n; j++ {
			fmt.Print(string(code[i].bits[j]), " ")
		}
		fmt.Println("")
	}
	// fmt.Println(tree)
	fmt.Println("Read the message, decoding")
	HuffmanDecode(tree)
}
