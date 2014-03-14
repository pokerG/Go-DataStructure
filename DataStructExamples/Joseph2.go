package main

import (
	"fmt"
)

type Memunit struct {
	id   int
	next *Memunit
	pre  *Memunit
}

var head, cur *Memunit

func run(m int) {
	for i := 2; i <= m; i++ {
		cur = cur.next
	}
	cur.pre.next = cur.next
	cur.next.pre = cur.pre
	fmt.Printf("No:%d is killed.\n", cur.id)
	cur = cur.next
}

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	head = &Memunit{}
	cur = head
	for i := 1; i <= n; i++ {
		cur.id = i
		if i != n {
			cur.next = &Memunit{}
		} else {
			cur.next = head
		}
		cur.next.pre = cur
		cur = cur.next
	}
	cur = head
	for cur.next != cur {
		run(m)
	}
	fmt.Printf("Winner is %d\n", cur.id)
	return
}
