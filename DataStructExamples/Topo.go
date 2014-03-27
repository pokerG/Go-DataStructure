package main

import (
	"fmt"
	"listGraph"
)

/*
test data:
10 11
0 9 1
7 0	1
8 7	1
8 5	1
3 4	1
3 1	1
3 2	1
4 1	1
4 2	1
5 4	1
5 6	1

*/
func main() {
	g := listGraph.NewGraph()
	g.Print()
	fmt.Println("TopoOrder:")
	g.TopoOrder()
	g.Print()
}
