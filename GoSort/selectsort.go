package GoSort

import ()

// SelectionSort
// O(n^2) running time.
func SelectSort(A []int) {
	l := len(A)
	if l < 2 {
		return
	}
	for i := 0; i < l; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if A[j] < A[min] {
				min = j
			}
		}
		if i != min {
			swap(&A[i], &A[min])
		}
	}
}
