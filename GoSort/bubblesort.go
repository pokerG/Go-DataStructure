//This package provides some sort way
//
//Copytright (C) 2014 by pokerG <pokerfacehlg@gmail.com>
package GoSort

import ()

//Bubble sort
// O(n^2) running time
func BubbleSort(A []int) {
	flag := true
	l := len(A)
	if l < 2 {
		return
	}
	for i := 0; i < l; i++ {
		flag = true
		for j := 0; j < l-1; j++ {
			if A[j] > A[j+1] {
				swap(&A[j], &A[j+1])
				flag = false
			}

		}
		if flag {
			break
		}
	}
}
