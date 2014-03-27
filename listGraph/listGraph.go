package listGraph

import (
	"Queue"
	"errors"
	"fmt"
	"reflect"
)

type vertexData int

type node struct { //out-edge form
	adjvex int
	cost   int
	next   *node
}

type vertexhead struct { //vertex form
	vertex    vertexData
	firstedge *node
}

type Graph struct { //Graph's list struct
	headlist []vertexhead
	n, e     int
}

var act []int    // use in Crucialpath
var ve, vl []int // use in Crucialpath
var size int

func NewGraph() *Graph {
	g := &Graph{}
	var tail, head, weight int
	fmt.Println("Please input the number of node and edge(n e)")
	fmt.Scanf("%d %d", &g.n, &g.e)
	g.headlist = make([]vertexhead, g.n)

	fmt.Println("Please input the edge(x y w)")
	for i := 0; i < g.e; i++ {
		fmt.Scanf("%d %d %d", &tail, &head, &weight)
		p := &node{}
		p.adjvex = head
		p.cost = weight
		p.next = g.headlist[tail].firstedge
		g.headlist[tail].firstedge = p
	}
	return g
}

func (a *Graph) Print() {
	for i := 0; i < a.n; i++ {
		fmt.Print(i, ": ")
		tmp := a.headlist[i].firstedge
		for tmp != nil {
			fmt.Print(tmp.adjvex, " ")
			tmp = tmp.next
		}
		fmt.Println("")

	}
}

func (a *Graph) TopoOrder() {
	q := Queue.NewGoQueueSize(a.n)
	q.MakeNull()
	// this := a.copy
	var inde []int
	inde = make([]int, a.n) //indegree
	for i := 0; i < a.n; i++ {
		inde[i] = indegree(i, *a)
		if inde[i] == 0 {
			q.Enter(i)
		}
	}
	// fmt.Println(inde)

	nodecount := 0
	for q.Empty() == nil {
		// q.Print()
		// fmt.Println("")
		t, _ := q.Front()
		b := reflect.ValueOf(t)
		v := int(b.Int())

		q.Delete()
		fmt.Print(v, " ")
		nodecount++

		tmp := a.headlist[v].firstedge
		for tmp != nil {
			inde[tmp.adjvex]--
			if inde[tmp.adjvex] == 0 {
				q.Enter(tmp.adjvex)
			}
			tmp = tmp.next
		}
	}
	if nodecount < 0 {
		fmt.Println("There is cycle on the digraph!")
	} else {
		fmt.Println("")
	}

}

func indegree(x int, g Graph) int {
	inde := 0
	for i := 0; i < g.n; i++ {
		tmp := g.headlist[i].firstedge
		for tmp != nil {
			if tmp.adjvex == x {
				inde++
			}
			tmp = tmp.next
		}
	}
	// fmt.Println(x, inde)
	return inde
}

func (a *Graph) Crucialpath() {
	act = make([]int, a.n*a.n)
	ve = make([]int, a.n)
	vl = make([]int, a.n)
	size = a.n
	for i := 0; i < a.n; i++ {
		tmp := a.headlist[i].firstedge
		for tmp != nil {
			act[i*a.n+tmp.adjvex] = tmp.cost
			tmp = tmp.next
		}
	}
	for i := 0; i < a.n; i++ {
		for j := 0; j < a.n; j++ {
			if act[i*a.n+j] == 0 {
				act[i*a.n+j] = -1
			}
			// fmt.Print(act[i*a.n+j], " ")
		}
		// fmt.Println("")
	}

	a.forward()
	a.backward()
	var l, e []int
	l = make([]int, a.n*a.n)
	e = make([]int, a.n*a.n)
	for i := 0; i < a.n; i++ {
		for j := 0; j < a.n; j++ {
			if act[i*a.n+j] > 0 {
				e[i*a.n+j] = ve[i]
				l[i*a.n+j] = vl[j] - act[i*a.n+j]
				act[i*a.n+j] = l[i*a.n+j] - e[i*a.n+j]
			}
		}
	}
	fmt.Println(l)
	fmt.Println(e)
	for i := 0; i < a.n; i++ {
		for j := 0; j < a.n; j++ {
			fmt.Print(act[i*a.n+j], " ")
		}
		fmt.Println("")
	}
	for i := 0; i < a.n; i++ {
		for j := 0; j < a.n; j++ {
			if act[i*a.n+j] >= 0 {
				fmt.Printf("From %d to %d : \t%d\t%d\t%d\n", i, j, e[i*a.n+j], l[i*a.n+j], act[i*a.n+j])
			}
		}
	}
	// for i := 0; i < a.n; i++ {
	// 	tmp := a.headlist[i].firstedge
	// 	if tmp != nil {
	// 		act[i*a.n+tmp.adjvex] = l(i, tmp.adjvex) - e(i, tmp.adjvex)
	// 		fmt.Printf("From %d to %d : \t%d\t%d\t%d\n", i, tmp.adjvex, e(i, tmp.adjvex), l(i, tmp.adjvex), act[i*a.n+tmp.adjvex])
	// 	}
	// }

}

func (a *Graph) forward() {
	q := Queue.NewGoQueueSize(a.n)
	q.MakeNull()
	// this := a.copy
	var inde []int
	inde = make([]int, a.n) //indegree
	for i := 0; i < a.n; i++ {
		inde[i] = indegree(i, *a)
		if inde[i] == 0 {
			q.Enter(i)
		}
	}
	// fmt.Println(inde)
	ve[1] = 0
	nodecount := 0
	for q.Empty() == nil {
		// q.Print()
		t, _ := q.Front()
		b := reflect.ValueOf(t)
		v := int(b.Int())

		q.Delete()
		/*for i := 0; i < a.n; i++ {
			tmp := a.headlist[i].firstedge
		flag:
			for tmp != nil {
				if tmp.adjvex == v {
					if act[i*a.n+v] > 0 && ve[v] < ve[i]+act[i*a.n+v] {
				ve[v] = ve[i] + act[i*a.n+v]
			}
					break flag
				}
				tmp = tmp.next
			}
		}*/

		for i := 0; i < a.n; i++ {
			if act[i*a.n+v] > 0 && ve[v] < ve[i]+act[i*a.n+v] {
				ve[v] = ve[i] + act[i*a.n+v]
			}
		}
		nodecount++

		tmp := a.headlist[v].firstedge
		for tmp != nil {
			inde[tmp.adjvex]--
			if inde[tmp.adjvex] == 0 {
				q.Enter(tmp.adjvex)
			}
			tmp = tmp.next
		}
	}
	if nodecount < 0 {
		panic(errors.New("There is cycle on the digraph!"))
	}
}

func (a *Graph) backward() {
	q := Queue.NewGoQueueSize(a.n)
	q.MakeNull()
	// this := a.copy
	b := a.reverse()
	// b.Print()

	var inde []int
	inde = make([]int, b.n) //indegree
	vl[b.n-1] = ve[b.n-1]
	for i := 0; i < b.n; i++ {
		inde[i] = indegree(i, *b)
		if inde[i] == 0 {
			q.Enter(i)
		}
	}
	// fmt.Println(inde)

	nodecount := 0
	for q.Empty() == nil {
		// q.Print()
		// fmt.Println("")
		t, _ := q.Front()
		x := reflect.ValueOf(t)
		v := int(x.Int())

		q.Delete()
		tmp := b.headlist[v].firstedge
		/*		for tmp != nil {
						i := tmp.adjvex
						if act[v*b.n+i] > 0 && vl[v] < vl[i]+act[v*b.n+i] {
					vl[v] = vl[i] - act[v*b.n+i]
				}
						tmp = tmp.next
					}
		*/
		for i := 0; i < b.n; i++ {
			if act[v*b.n+i] > 0 && vl[v] < vl[i]+act[v*b.n+i] {
				vl[v] = vl[i] - act[v*b.n+i]
			}
		}
		nodecount++

		tmp = b.headlist[v].firstedge
		for tmp != nil {
			inde[tmp.adjvex]--
			if inde[tmp.adjvex] == 0 {
				q.Enter(tmp.adjvex)
			}
			tmp = tmp.next
		}
	}
	if nodecount < 0 {
		panic(errors.New("There is cycle on the digraph!"))
	}
}

// func (a *Graph) copy() *Graph {
// 	this := &Graph{}
// 	this.n = a.n
// 	this.e = a.e
// 	this.headlist = make([]vertexhead, this.n)

// 	this.edge = make([]int, this.n*this.n)
// 	for i, v := range a.edge {
// 		this.edge[i] = v
// 	}
// 	return this
// }

// func e(k, m int) int {
// 	return ve[k]
// }
// func l(k, m int) int {
// 	return vl[m] - act[k*size+m]
// }

func (a *Graph) reverse() *Graph { //tail and head reverse
	b := &Graph{}
	b.n = a.n
	b.e = a.e
	b.headlist = make([]vertexhead, b.n)
	for i := 0; i < a.n; i++ {
		tmp := a.headlist[i].firstedge
		for tmp != nil {
			p := &node{}
			p.adjvex = i
			p.cost = tmp.cost
			p.next = b.headlist[tmp.adjvex].firstedge
			b.headlist[tmp.adjvex].firstedge = p
			tmp = tmp.next
		}
	}
	return b
}
