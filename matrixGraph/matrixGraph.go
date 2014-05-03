//This package achieve graph's create (use array struct) and some traversal algorithm
//inlude Kruskal Prim Dijkstra Floyd
//Dijstra and Prim use max heap to optimize
//
//Copytright (C) 2014 by pokerG <pokerfacehlg@gmail.com>
package matrixGraph

import (
	"container/heap"
	"fmt"
	// "sort"
)

const maxSize = 100
const infinity = 1 << 30

type Graph struct {
	n, e int //number of node and edge
	edge []int
}

type edges struct {
	begin, end int
	cost       int
	inHeap     bool //use in prioityQueue
}

var Edges []edges

//NewGraph creates a new Graph
func NewGraph() *Graph {
	g := &Graph{}
	fmt.Println("Please input the number of node and edge(n e)")
	fmt.Scanf("%d %d\n", &g.n, &g.e)
	fmt.Println("Please input the edge(x y w)")
	g.edge = make([]int, g.n*g.n)
	Edges = make([]edges, g.e)

	for i := 0; i < g.e; i++ {
		var x, y, w int
		fmt.Scanf("%d %d %d\n", &x, &y, &w)

		Edges[i].begin = x
		Edges[i].end = y
		Edges[i].cost = w
		Edges[i].inHeap = false // -false is not in heap

		g.edge[x*g.n+y] = w
	}
	return g
}

//Print the adjacency matrix
func (this *Graph) Print() {
	fmt.Printf("The Graph have %d nodes and %d edges\n", this.n, this.e)
	for i := 0; i < this.n; i++ {
		for j := 0; j < this.n; j++ {
			fmt.Print(this.edge[i*this.n+j], " ")
		}
		fmt.Println("")
	}
}

//Kruskal
//use Disjoint Set
func (a *Graph) Kruskal() {
	var bnf, edf int
	var parent [maxSize]int
	this := a.copy()
	e := newEdges(this)
	sort(e)
	for i := 0; i < this.n; i++ {
		parent[i] = i
	}
	for i := 0; i < this.e; i++ {
		bnf = find(e[i].begin, &parent)
		edf = find(e[i].end, &parent)
		if bnf != edf {
			parent[bnf] = edf
			fmt.Println("(", e[i].begin, ",", e[i].end, ",", e[i].cost, ")")
		}
	}

}

func (a *Graph) Prim() {
	var lowcost, closet []int
	var i, j, k int
	var min int
	this := a.copy()
	lowcost = make([]int, this.n)
	closet = make([]int, this.n)
	// change to undirected graph
	// 0 convert to infinity
	for i, _ := range this.edge {
		if this.edge[i] != 0 {
			this.edge[i%this.n*this.n+i/this.n] = this.edge[i]
		}
		if this.edge[i] == 0 {
			this.edge[i] = infinity
		}
	}

	for i = 1; i < this.n; i++ {
		lowcost[i] = this.edge[i]
		closet[i] = 0
	}
	for i = 1; i < this.n; i++ {
		min = lowcost[i]
		k = i
		for j = 1; j < this.n; j++ {
			if lowcost[j] < min {
				min = lowcost[j]
				k = j
			}
		}
		fmt.Println("(", k, ",", closet[k], ",", min, ")")
		lowcost[k] = infinity
		this.edge[k*this.n+closet[k]] = infinity
		this.edge[closet[k]*this.n+k] = infinity
		for j = 0; j < this.n; j++ {
			if this.edge[k*this.n+j] < lowcost[j] {
				lowcost[j] = this.edge[k*this.n+j]
				closet[j] = k
			}
		}
	}
}

//Use max heap to optimize
func (a *Graph) PrimHeap() {
	this := a.copy()
	// change to undirected graph
	// 0 convert to infinity
	for i, _ := range this.edge {
		if this.edge[i] != 0 {
			this.edge[i%this.n*this.n+i/this.n] = this.edge[i]
		}
		if this.edge[i] == 0 {
			this.edge[i] = infinity
		}
	}

	h := &priorityQueue{}
	heap.Init(h)

	var vis []bool
	vis = make([]bool, this.n)
	vis[0] = true

	var pre int // record the begin node

	for i := 0; i < this.n; i++ {
		if this.edge[i] != infinity {
			heap.Push(h, &space{i, this.edge[i], -1})
		}
	}
	pre = 0
	for i := 1; i < this.n; i++ {
		e := heap.Pop(h).(*space)
		v := e.end
		vis[v] = true
		fmt.Println("(", pre, ",", v, ",", e.cost, ")")
		pre = v
		for i := 0; i < this.n; i++ {
			if this.edge[v*this.n+i] != infinity && !vis[i] {
				heap.Push(h, &space{i, this.edge[v*this.n+i], -1})
			}
		}
	}
}

//Use in Kruskal
func newEdges(g *Graph) []edges {
	var e []edges
	e = make([]edges, g.e)
	n := 0
	for i := 0; i < g.n; i++ {
		for j := 0; j < g.n; j++ {
			if g.edge[i*g.n+j] != 0 {
				e[n].begin = i
				e[n].end = j
				e[n].cost = g.edge[i*g.n+j]
				n++

			}
		}
	}
	return e
}

//Use in Kruskal
func sort(e []edges) {
	for j := 0; j < len(e)-1; j++ {
		for i := 0; i < len(e)-1-j; i++ {
			if e[i].cost > e[i+1].cost {
				temp := e[i]
				e[i] = e[i+1]
				e[i+1] = temp
			}
		}
	}
}

//Use in Kruskal
func find(x int, p *[maxSize]int) int {
	if x == p[x] {
		return x
	}
	p[x] = find(p[x], p)
	return p[x]
}

func (a *Graph) Dijkstra() []int {
	var k int
	var used []bool
	var dis []int
	this := a.copy()
	// fmt.Println(&a, &this)
	used = make([]bool, this.n)
	dis = make([]int, this.n)

	// 0 convert to infinity
	for i, _ := range this.edge {
		if this.edge[i] == 0 {
			this.edge[i] = infinity
		}
	}
	for i := 0; i < this.n; i++ {
		dis[i] = this.edge[i]
	}
	used[0] = true

	for i := 0; i < this.n-1; i++ {
		tmin := infinity
		for j := 1; j < this.n; j++ {
			if !used[j] && tmin > dis[j] {
				tmin = dis[j]
				k = j
			}
		}
		used[k] = true
		for j := 1; j < this.n; j++ {
			if used[j] {
				continue
			}
			if dis[k]+this.edge[k*this.n+j] < dis[j] {
				dis[j] = dis[k] + this.edge[k*this.n+j]
			}
		}
	}
	// a.Print()
	// this.Print()
	return dis
}

//Use max heap to optimize
func (a *Graph) DijkstraHeap() []int {
	this := a.copy()
	// fmt.Println(&a, &this)

	dis := make([]int, this.n)

	h := &priorityQueue{}
	heap.Init(h)

	for _, v := range Edges {
		if v.begin == 0 {
			v.inHeap = true
			heap.Push(h, &space{v.end, v.cost, -1})
		}
	}

	for i := 0; i < this.n; i++ {
		e := heap.Pop(h).(*space)

		v := e.end

		dis[v] = e.cost
		for i, _ := range Edges {
			if Edges[i].begin == v {
				if Edges[i].inHeap {
					for i, _ := range *h {
						if (*h)[i].end == Edges[i].end {
							h.update((*h)[i], (*h)[i].end, e.cost+Edges[i].cost)
						}
					}
				} else {
					heap.Push(h, &space{Edges[i].end, e.cost + Edges[i].cost, -1})
				}
			}
		}

	}
	return dis

}

func (a *Graph) Floyd() []int {
	var d, p []int
	this := a.copy()
	d = make([]int, this.n*this.n)
	p = make([]int, this.n*this.n)

	// 0 convert to infinity
	for i, _ := range this.edge {
		if this.edge[i] == 0 {
			this.edge[i] = infinity
		}
	}

	for i, v := range this.edge {
		d[i] = v
		p[i] = -1
	}
	for k := 0; k < this.n; k++ {
		for i := 0; i < this.n; i++ {

			aki := d[i*this.n+k]
			if aki == infinity { // ignore the path of not exist
				continue
			}

			for j := 0; j < this.n; j++ {
				if d[i*this.n+k]+d[k*this.n+j] < d[i*this.n+j] {
					d[i*this.n+j] = d[i*this.n+k] + d[k*this.n+j]
					p[i*this.n+j] = k
				}
			}
		}
	}
	return d
}

//prevent destroy the primary data
func (a *Graph) copy() *Graph {
	this := &Graph{}
	this.n = a.n
	this.e = a.e
	this.edge = make([]int, this.n*this.n)
	for i, v := range a.edge {
		this.edge[i] = v
	}
	return this
}
