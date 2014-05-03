//This package provides some sort way
//
//Copytright (C) 2014 by pokerG <pokerfacehlg@gmail.com>
package GoSort

import (
	"runtime"
)

//Merge sort parallel
func MergeSortP(A []int) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	n = len(A)
	s := make(chan bool)

	go msp(A, s)
	<-s

}

func msp(A []int, s chan bool) {
	l := len(A)
	if l < 2 {
		s <- true
		return
	}
	if l < 7 {
		InsertSort(A)
		s <- true
		return
	}

	middle := l / 2
	left := A[:middle]
	right := A[middle:]
	s1 := make(chan bool)
	s2 := make(chan bool)
	go msp(left, s1)
	go msp(right, s2)
	<-s1
	<-s2
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
	s <- true
}
