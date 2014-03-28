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

7 9
0 2 1
2 1 1
1 0 1
2 3 1
3 4 1
4 3 1
5 6 1
6 5 1
6 4 1
*/

func main() {
	g := listGraph.NewGraph()
	g.Korasaju()
}
