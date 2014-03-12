package main

import (
	"fmt"
)

type Triple struct {
	i, j int
	data int
}

func MatrixZip(arr [][]int) []Triple {
	tri := make([]Triple, 1)
	p := 0
	tri[0].i = len(arr)
	tri[0].j = len(arr[0])
	for i, _ := range arr {
		for j, v := range arr[i] {
			if arr[i][j] != 0 {
				tri = append(tri, Triple{i, j, v})
				p++
			}
		}
	}
	tri[0].data = p
	return tri
}

func MatrixUZip(tri []Triple) [][]int {
	var arr [][]int
	arr = make([][]int, 0)
	p := 1
	for i := 0; i < tri[0].i; i++ {
		arr = append(arr, make([]int, tri[0].j))
		for j := 0; j < tri[0].j; j++ {
			if p <= tri[0].data && i == tri[p].i && j == tri[p].j {
				arr[i][j] = tri[p].data
				p++
			} else {
				arr[i][j] = 0
			}
		}
	}
	return arr
}

func TransposeSMatrix(triA []Triple) []Triple {
	triB := make([]Triple, 1)
	triB[0].i = triA[0].j
	triB[0].j = triA[0].i
	triB[0].data = triA[0].data
	if triB[0].data > 0 {
		for col := 0; col < triA[0].j; col++ {
			for p := 1; p <= triA[0].data; p++ {
				if triA[p].j == col {
					triB = append(triB, Triple{triA[p].j, triA[p].i, triA[p].data})
				}
			}
		}
	}
	return triB
}

func main() {
	var x, y int
	fmt.Print("input the size of matrix(x,y): ")
	fmt.Scanf("%d %d", &x, &y)
	var arr [][]int
	arr = make([][]int, 0)
	// fmt.Println(x, y)
	for i := 0; i < x; i++ {
		arr = append(arr, make([]int, y))
		for j := 0; j < y; j++ {
			fmt.Scanf("%d", &arr[i][j])
		}
	}
	triA := MatrixZip(arr)

	fmt.Println(triA)
	triB := TransposeSMatrix(triA)
	a := MatrixUZip(triB)
	fmt.Println(a)
}
