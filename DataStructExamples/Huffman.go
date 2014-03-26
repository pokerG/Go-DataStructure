package main

import (
	"fmt"
)

/*
test data
a 0.15
b 0.45
c 0.25
d 0.20
e 0.08
*/
const n int = 5         //the number of leaf
const m int = (2*n - 1) //the number of node
const maxVal float64 = 10000.0
const maxSize int = 100 //huffman code's wightest digit

type HufmTree struct {
	ch                     string
	weight                 float64
	lchild, rchild, parent int
}

type CodeType struct {
	bits  [n]byte //bit string
	start int     //code's start position in bits
	ch    string
}

func Huffman(tree []HufmTree) {
	var i, j, p1, p2 int
	var small1, small2, f float64
	var c string
	for i = 0; i < m; i++ {
		tree[i].parent = 0
		tree[i].lchild = -1
		tree[i].rchild = -1
		tree[i].weight = 0.0
	}
	fmt.Printf("Please input %d node's byte and weight(sepatate by space)\n", n)
	for i = 0; i < n; i++ {
		fmt.Printf("Input the %dth byte and weight: ", i+1)
		fmt.Scanf("%v %v\n", &c, &f)
		tree[i].ch = c
		tree[i].weight = f
	}
	for i = n; i < m; i++ {
		p1 = 0
		p2 = 0
		small1 = maxVal
		small2 = maxVal
		for j = 0; j < i; j++ {
			if tree[j].parent == 0 {
				if tree[j].weight < small1 {
					small2 = small1
					small1 = tree[j].weight
					p2 = p1
					p1 = j
				} else if tree[j].weight < small2 {
					small2 = tree[j].weight
					p2 = j
				}
			}
		}
		tree[p1].parent = i
		tree[p2].parent = i
		tree[i].lchild = p1
		tree[i].rchild = p2
		tree[i].weight = tree[p1].weight + tree[p2].weight
	}
}

func HuffmanCode(code []CodeType, tree []HufmTree) {
	var i, c, p int
	var cd CodeType //buffer
	for i = 0; i < n; i++ {
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

func HuffmanDecode(tree []HufmTree) {
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
	fmt.Printf("A total of %d characters\n", n)
	tree := make([]HufmTree, m)
	code := make([]CodeType, n)

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
