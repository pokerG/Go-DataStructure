package main

import (
	// "fmt"
	"listGraph"
)

/*
test data:
9 1
10 1 6
0 2 4
0 3 5
1 4 1
2 4 1
3 5 2
4 6 9
4 7 7
5 7 4
6 8 2
7 8 4

*/
func main() {
	g := listGraph.NewGraph()
	g.Crucialpath()
}
