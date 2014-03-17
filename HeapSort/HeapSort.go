package HeapSort

import (
	// "fmt"
	"math"
)

var HeapSize int

func Parent(i int) int {
	return int(math.Ceil(float64(i) / 2))
}

func Left(i int) int {
	return 2 * i
}

func Right(i int) int {
	return 2*i + 1
}

func swap(a, b *int) {
	c := *a
	(*a) = (*b)
	(*b) = c
}

func MaxHeapIFY(A []int, i int) {
	var largest int
	l := Left(i)
	r := Right(i)
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
		MaxHeapIFY(A, largest)
	}
}

func BuildMaxHeap(A []int) {
	HeapSize = len(A) - 1
	for i := int(math.Ceil(float64(len(A)) / 2)); i >= 1; i-- {
		MaxHeapIFY(A, i)
	}
}

func HeapSort(A []int) {
	BuildMaxHeap(A)
	for i := len(A) - 1; i >= 2; i-- {
		swap(&A[1], &A[i])
		HeapSize--
		MaxHeapIFY(A, 1)

	}
}
