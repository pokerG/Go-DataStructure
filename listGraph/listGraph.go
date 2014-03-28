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
	vertex    vertexData //not use
	firstedge *node
}

type Graph struct { //Graph's list struct
	headlist []vertexhead
	n, e     int
}

var act []int    // use in Crucialpath
var ve, vl []int // use in Crucialpath

var count int //use in FindArticul
var dnf []int //use in FindArticul
var low []int //use in FindArticul
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

func (a *Graph) deleteDirect() *Graph {
	b := &Graph{}
	b.n = a.n
	b.e = a.e
	b.headlist = make([]vertexhead, b.n)
	// edge := make([]int, b.n*b.n)
	// for i := 0; i < a.n; i++ {
	// 	tmp := a.headlist[i].firstedge
	// 	for tmp != nil {
	// 		edge[i*a.n+tmp.adjvex] = tmp.cost
	// 		edge[tmp.adjvex*a.n+i] = tmp.cost
	// 		tmp = tmp.next
	// 	}
	// }
	// // fmt.Println("!!")
	// for i := 0; i < b.n; i++ {
	// 	for j := 0; j < b.n; j++ {
	// 		// fmt.Println(i*b.n+j, len(edge))
	// 		if edge[i*b.n+j] > 0 {
	// 			tmp := b.headlist[i].firstedge
	// 			if tmp == nil {
	// 				p := &node{}
	// 				p.adjvex = j
	// 				p.cost = edge[i*b.n+j]
	// 				p.next = nil
	// 				b.headlist[i].firstedge = p
	// 			} else {
	// 				p := tmp
	// 				for tmp != nil {
	// 					p = tmp
	// 					tmp = tmp.next
	// 				}
	// 				tmp = &node{}
	// 				tmp.adjvex = j
	// 				tmp.next = nil
	// 				p.next = tmp
	// 			}
	// 		}
	// 	}
	// }
	for i := 0; i < a.n; i++ {
		tmp := a.headlist[i].firstedge
		for tmp != nil {
			p := &node{}
			p.adjvex = tmp.adjvex
			p.cost = tmp.cost
			p.next = b.headlist[i].firstedge
			b.headlist[i].firstedge = p
			tmp2 := b.headlist[tmp.adjvex].firstedge
			for tmp2 != nil {
				if tmp2.adjvex == i {
					goto flag
				}
				tmp2 = tmp2.next
			}
			p = &node{}
			p.adjvex = i
			p.cost = tmp.cost
			p.next = b.headlist[tmp.adjvex].firstedge
			b.headlist[tmp.adjvex].firstedge = p
		flag:
			tmp = tmp.next
		}
	}
	return b
}

func (a *Graph) FindArticul() {
	b := a.deleteDirect()
	b.Print()
	count = 1
	dnf = make([]int, b.n)
	low = make([]int, b.n)
	dnf[0] = 1

	p := b.headlist[0].firstedge
	v := p.adjvex

	//a.dfsArticul(v) fuck this!!!!!!
	b.dfsArticul(v)
	if count < b.n {
		fmt.Print("0 ")
		for p.next != nil {
			p = p.next
			v = p.adjvex
			if dnf[v] == 0 {
				b.dfsArticul(v)
			}
		}
	}
}

func (a *Graph) dfsArticul(v int) {
	count++
	min := count
	dnf[v] = min
	// fmt.Println("!!!", v)
	for p := a.headlist[v].firstedge; p != nil; p = p.next {
		w := p.adjvex
		// fmt.Println("###", w)
		if dnf[w] == 0 {
			a.dfsArticul(w)
			if low[w] < min {
				min = low[w]
			}
			if low[w] >= dnf[v] {
				// fmt.Print("!!!!!")
				fmt.Print(v, " ")
			}

		} else if dnf[w] < min {
			min = dnf[w]
		}
	}
	low[v] = min
}

var visited []bool
var inOrder []int

// var count int
func (a *Graph) Korasaju() {
	k := 1
	visited = make([]bool, a.n)
	inOrder = make([]int, a.n)
	for i := 0; i < a.n; i++ {
		visited[i] = false
	}
	count = 0
	for i := 0; i < a.n; i++ {
		if !visited[i] {
			dfs(a, i)
		}
	}
	for i := 0; i < a.n; i++ {
		visited[i] = false
	}
	b := a.reverse()
	for i := b.n - 1; i >= 0; i-- {
		v := inOrder[i]
		if !visited[v] {
			fmt.Printf("\nThe %dth connected component vertex: ", k)
			k++
			revDfs(b, v)
			fmt.Println("")
		}

	}
}

func dfs(g *Graph, x int) {
	visited[x] = true
	for p := g.headlist[x].firstedge; p != nil; p = p.next {
		if !visited[p.adjvex] {
			dfs(g, p.adjvex)
		}
	}
	inOrder[count] = x
	count++
}

func revDfs(g *Graph, x int) {
	visited[x] = true
	fmt.Print(x, " ")
	for p := g.headlist[x].firstedge; p != nil; p = p.next {
		if !visited[p.adjvex] {
			revDfs(g, p.adjvex)
		}
	}
}
