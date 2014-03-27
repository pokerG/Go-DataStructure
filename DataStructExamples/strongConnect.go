package main

import (
	"listGraph"
	// "fmt"
)

/*
test data:
6 8
0 1 1
0 2 1
3 0 1
2 3 1
2 5 1
5 4 1
4 1 1
1 5 1
*/

func main() {
	g := listGraph.NewGraph()
	g.Korasaju()
}
