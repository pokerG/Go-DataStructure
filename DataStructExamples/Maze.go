package main

import (
	"Stack"
	"fmt"
	"reflect"
)

/*
test data:
4,4
1 1 1 0
1 0 1 1
1 1 0 0
0 1 1 2

4,4
1 1 1 1
1 0 1 1
1 1 0 0
0 1 0 2
2 is export
*/
type Map struct {
	i, j int
}

var x, y int
var maze [][]int
var visited [][]bool

var w = []int{1, 0, -1, 0}
var h = []int{0, 1, 0, -1}

func main() {
	fmt.Println("Please input the size of maze(x,y): ")
	fmt.Scanf("%d,%d", &x, &y)
	maze = make([][]int, 0)
	maze = append(maze, make([]int, y+2))
	visited = make([][]bool, 0)
	visited = append(visited, make([]bool, y+2))
	for i := 1; i <= x; i++ {
		visited = append(visited, make([]bool, y+2))
		maze = append(maze, make([]int, y+2))
		for j := 1; j <= y; j++ {
			fmt.Scanf("%d", &maze[i][j])
		}
	}
	maze = append(maze, make([]int, y+2))
	visited = append(visited, make([]bool, y+2))
	for i := 0; i <= y+1; i++ {
		maze[0][i] = 0
	}
	for j := 0; j <= x+1; j++ {
		maze[j][0] = 0
	}
	fmt.Println(maze)
	if DFS(1, 1) {
		fmt.Println("You can flee")
	} else {
		fmt.Println("You can't flee")
	}
}

func DFS(i, j int) bool {
	s := Stack.NewGoStack()
	s.Push(Map{i, j})
	visited[i][j] = true
	for s.Empty() == nil {
		m, _ := s.Top()
		t := reflect.ValueOf(m)

		v := Map{int(t.Field(0).Int()), int(t.Field(1).Int())}
		// fmt.Println(v.i, v.j)
		if maze[v.i][v.j] == 2 {
			return true
		}
		// s.Print()
		if !visited[v.i+w[0]][v.j+h[0]] && maze[v.i+w[0]][v.j+h[0]] != 0 {
			visited[v.i+w[0]][v.j+h[0]] = true
			s.Push(Map{v.i + w[0], v.j + h[0]})
			continue
		}
		// fmt.Println(v.i+w[1], v.j+h[1])
		if !visited[v.i+w[1]][v.j+h[1]] && maze[v.i+w[1]][v.j+h[1]] != 0 {
			visited[v.i+w[1]][v.j+h[1]] = true
			s.Push(Map{v.i + w[1], v.j + h[1]})
			continue
		}
		if !visited[v.i+w[2]][v.j+h[2]] && maze[v.i+w[2]][v.j+h[2]] != 0 {
			visited[v.i+w[2]][v.j+h[2]] = true
			s.Push(Map{v.i + w[2], v.j + h[2]})
			continue
		}
		if !visited[v.i+w[3]][v.j+h[3]] && maze[v.i+w[3]][v.j+h[3]] != 0 {
			visited[v.i+w[3]][v.j+h[3]] = true
			s.Push(Map{v.i + w[3], v.j + h[3]})
			continue
		}
		s.Pop()
	}
	return false
}
