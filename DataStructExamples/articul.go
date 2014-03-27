package main

import (
	"listGraph"
	// "fmt"
)

/*
test data:
7 10
0 1 1
0 2 1
0 3 1
0 4 1
1 3 1
1 4 1
3 4 1
2 5 1
2 6 1
5 6 1

7 7
0 1 1
0 2 1
1 3 1
3 4 1
2 5 1
2 6 1
5 6 1
*/

func main() {
	g := listGraph.NewGraph()
	g.FindArticul()
}
