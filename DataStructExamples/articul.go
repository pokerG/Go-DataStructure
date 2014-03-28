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

7 21
0 1
0 3
1 0
1 2
1 3
1 4
2 1
2 3
2 4
2 5
3 0
3 1
3 2
3 4
4 1
4 2
4 3
4 5
5 4
5 6
6 5
*/

func main() {
	g := listGraph.NewGraph()
	g.FindArticul()
}
