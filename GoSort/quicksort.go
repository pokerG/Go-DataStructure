package GoSort

import (
// "fmt"
)

func QuickSort(A []int) {
	l := len(A)
	val := A[0]
	// fmt.Println(l)
	if l < 2 {
		return
	}
	i := 0
	j := l - 1
	// fmt.Println("i:", i, " j: ", j, A)
	for i < j {
		for ; j > i; j-- {
			if A[j] <= val {
				swap(&A[i], &A[j])
				break
			}
		}
		for ; i < j; i++ {
			if A[i] > val {
				swap(&A[i], &A[j])
				break
			}
		}
	}
	// A[i] = val
	// fmt.Println("i:", i, " j: ", j, A)
	if i > 1 {
		QuickSort(A[:i])
	}
	if i < l-1 {
		QuickSort(A[i+1:])
	}

}
