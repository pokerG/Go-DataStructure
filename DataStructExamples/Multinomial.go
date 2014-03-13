package main

import (
	"LinkedList"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	var m1, m2 string
	fmt.Println("Please input the first multinomial")
	fmt.Scanf("%s\n", &m1)
	fmt.Println("Please input the second multinomial")
	fmt.Scanf("%s\n", &m2)
	l1 := EquationToList(m1)
	l2 := EquationToList(m2)
	l1.Print()
	l2.Print()
	menu()
	var index string
	for {
		fmt.Scanf("%s\n", &index)
		switch index {
		case "+":
			l := PolyAdd(l1, l2)
			l.Print()
		case "-":
			l := PolySub(l1, l2)
			l.Print()
		case "*":
			l := PolyMulti(l1, l2)
			// l = PolyAdd(l, l)
			l.Print()
		case "0":
			return
		}
	}
}

func menu() {
	fmt.Println("Please input + - *")
	fmt.Sprintln("0 is Exit")
}

func EquationToList(m string) *LinkedList.LinkedList {
	l := LinkedList.NewLinkedList()
	fmt.Println(m)
	sg := strings.Split(m, "+")
f:
	for _, v := range sg {
		var node [2]int
		var err error
		if !strings.Contains(v, "x") {
			node[0], _ = strconv.Atoi(v)
			node[1] = 0
			n := LinkedList.NewNode(node, nil)
			l.Insert(*n)
			continue f
		}
		tmp := strings.Split(v, "x")
		node[0], err = strconv.Atoi(tmp[0])
		if err != nil {
			node[0] = 1
		}
		if strings.Contains(tmp[1], "^") {
			node[1], _ = strconv.Atoi(string([]byte(tmp[1])[1:]))
		} else {
			node[1] = 1
		}
		n := LinkedList.NewNode(node, nil)
		l.Insert(*n)
	}
	return l
}

func PolyAdd(l1, l2 *LinkedList.LinkedList) *LinkedList.LinkedList {
	l := LinkedList.NewLinkedList()
	n1 := l1.Head
	n2 := l2.Head
	for n1 != nil || n2 != nil {
		if n1 != nil && (n2 == nil || getData(n1, 1) > getData(n2, 1)) {
			n := LinkedList.NewNode(n1.Data, nil)
			l.Insert(*n)
			n1 = n1.Next
		} else if n2 != nil && (n1 == nil || getData(n1, 1) < getData(n2, 1)) {
			n := LinkedList.NewNode(n2.Data, nil)
			l.Insert(*n)
			n2 = n2.Next
		} else if getData(n1, 1) == getData(n2, 1) {
			var tmp [2]int
			tmp[0] = getData(n1, 0) + getData(n2, 0)
			tmp[1] = getData(n2, 1)
			n := LinkedList.NewNode(tmp, nil)
			l.Insert(*n)
			n1 = n1.Next
			n2 = n2.Next
		}
	}
	return l
}

func PolySub(l1, l2 *LinkedList.LinkedList) *LinkedList.LinkedList {
	l := LinkedList.NewLinkedList()
	n1 := l1.Head
	n2 := l2.Head
	for n1 != nil || n2 != nil {
		if n2 == nil || getData(n1, 1) > getData(n2, 1) {
			n := LinkedList.NewNode(n1.Data, nil)
			l.Insert(*n)
			n1 = n1.Next
		} else if n1 == nil || getData(n1, 1) < getData(n2, 1) {
			n := LinkedList.NewNode(n2.Data, nil)
			l.Insert(*n)
			n2 = n2.Next
		} else if getData(n1, 1) == getData(n2, 1) {
			var tmp [2]int
			tmp[0] = getData(n1, 0) - getData(n2, 0)
			tmp[1] = getData(n2, 1)
			n := LinkedList.NewNode(tmp, nil)
			l.Insert(*n)
			n1 = n1.Next
			n2 = n2.Next
		}
	}
	return l
}

func getData(d *LinkedList.Node, n int) int {
	v := reflect.ValueOf(d.Data)
	b := int(v.Index(n).Int())
	return b
}

func MultOneTerm(n *LinkedList.Node, l *LinkedList.LinkedList) *LinkedList.LinkedList {
	var newn [2]int
	n0 := getData(n, 0)
	n1 := getData(n, 1)
	newl := LinkedList.NewLinkedList()
	ln := l.Head
	for ln != nil {
		ln0 := getData(ln, 0)
		ln1 := getData(ln, 1)
		newn[0] = n0 * ln0
		newn[1] = n1 + ln1
		newl.Insert(*LinkedList.NewNode(newn, nil))
		ln = ln.Next
	}
	return newl
}

func PolyMulti(l1, l2 *LinkedList.LinkedList) *LinkedList.LinkedList {
	n := l1.Head
	l := LinkedList.NewLinkedList()
	for ; n != nil; n = n.Next {
		h := MultOneTerm(n, l2)
		l = PolyAdd(l, h)
	}
	return l
}
