package GoSort

import ()

// InsertionSort  with O(n^2) running time.
func InsertSort(A []int) {
	l := len(A)
	if l < 2 {
		return
	}
	for i := 1; i < l; i++ {
		pivot := A[i]
		j := i
		for ; j > 0 && pivot < A[j-1]; j-- {
			A[j] = A[j-1]
		}
		A[j] = pivot
	}
}
