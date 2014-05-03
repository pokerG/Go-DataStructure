package main

import (
	// "bufio"
	"fmt"
	// "os"
	// "strings"
	"math"
	"matrixGraph"
	// "sort"
)

/*
test data:
5 6
0 1 3
0 2 4
2 3 5
3 4 1
2 4 2
0 3 9

0 1 1
0 2 2
0 3 3
2 3 1
2 4 4
3 4 2

*/

type c struct {
	a, b int
}

func main() {
	g := matrixGraph.NewGraph()
	g.Print()

	fmt.Println("Kruskal:")
	g.Kruskal()

	fmt.Println("Prim:")
	g.Prim()
	dis := g.Dijkstra()
	fmt.Println(dis)
	p := g.Floyd()
	l := int(math.Sqrt(float64(len(p))))
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			fmt.Print(p[i*l+j], " ")
		}
		fmt.Println("")
	}
	// fmt.Println(p)
	// var a [][]int
	// a = make([10][10]int, 10)

	// a[2][1] = 1
	// fmt.Println(a)
	// a := []int{1, 3, 2, 4, 5}
	// sort.Ints(a)
	// fmt.Println(a)
	// t1 := &c{1, 2}
	// t2 := t1
	// fmt.Println(&t1, &t2)
	// t2.a = 2
	// fmt.Println(t1, t2)
}
