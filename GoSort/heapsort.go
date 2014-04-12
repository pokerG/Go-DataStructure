package GoSort

import (
	// "fmt"
	"math"
)

// Binary heap sort
// O(n*logn) running time with constant extra space.
var HeapSize int

func parent(i int) int {
	return int(math.Ceil(float64(i) / 2))
}

func left(i int) int {
	return 2 * i
}

func right(i int) int {
	return 2*i + 1
}

func maxHeapIFY(A []int, i int) {
	var largest int
	l := left(i)
	r := right(i)
	// fmt.Println(l, r)
	if l <= HeapSize && A[l] > A[i] {
		largest = l
	} else {
		largest = i
	}
	if r <= HeapSize && A[r] > A[largest] {
		largest = r
	}
	if largest != i {
		swap(&A[i], &A[largest])
		maxHeapIFY(A, largest)
	}
}

func buildMaxHeap(A []int) {
	HeapSize = len(A) - 1
	for i := int(math.Ceil(float64(len(A)) / 2)); i >= 1; i-- {
		maxHeapIFY(A, i)
	}
}

func HeapSort(A []int) {
	// b := append(A, 0)
	// swap(&b[0], &b[len(b)-1])
	// copy(A, b)
	buildMaxHeap(A)
	for i := len(A) - 1; i >= 2; i-- {
		swap(&A[1], &A[i])
		HeapSize--
		maxHeapIFY(A, 1)

	}
}
