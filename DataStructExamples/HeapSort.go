package main

import "fmt"
import (
	"HeapSort"
	// "math"
)

// import "strconv"
// import "reflect"

var A []int

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	A = make([]int, n+1)
	A[0] = 0
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &A[i])
	}
	HeapSort.HeapSort(A)
	fmt.Println(A)
}
