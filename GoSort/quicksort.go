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
		for j > i && A[j] >= val {
			j--
		}
		if i < j {
			A[i] = A[j]
		}
		for i < j && A[i] <= val {
			i++
		}
		if i < j {
			A[j] = A[i]
		}
	}
	A[i] = val
	// A[i] = val
	// fmt.Println("i:", i, " j: ", j, A)
	if i > 1 {
		QuickSort(A[:i])
	}
	if i < l-1 {
		QuickSort(A[i+1:])
	}

}
