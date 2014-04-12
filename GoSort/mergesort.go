package GoSort

import ()

// MergeSort will sort the given slice of strings using the
// basic merge sort algorithm, with O(n log n) running time.
func MergeSort(A []int) {
	l := len(A)
	if l < 2 {
		return
	}
	if l < 7 {
		InsertSort(A)
		return
	}

	middle := l / 2
	left := A[:middle]
	right := A[middle:]
	MergeSort(left)
	MergeSort(right)

	result := make([]int, 0, l)
	li := 0
	ll := len(left)
	ri := 0
	rl := len(right)

	for li < ll && ri < rl {
		if left[li] <= right[ri] {
			result = append(result, left[li])
			li++
		} else {
			result = append(result, right[ri])
			ri++
		}
	}
	if li < ll {
		result = append(result, left[li:]...)
	} else if ri < rl {
		result = append(result, right[ri:]...)
	}
	copy(A, result)
}
