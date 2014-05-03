//This package provides some sort way
//
//Copytright (C) 2014 by pokerG <pokerfacehlg@gmail.com>
package GoSort

import ()

// ShellSort
// O(n^2) though often performs better in practice.
func ShellSort(A []int) {
	l := len(A)
	if l < 2 {
		return
	}
	inc := l / 2
	for inc > 0 {
		for i := inc; i < l; i++ {
			temp := A[i]
			j := i
			for j >= inc && A[j-inc] > temp {
				A[j] = A[j-inc]
				j -= inc
			}
			A[j] = temp
		}
		if inc == 2 {
			inc = 1
		} else {
			inc = inc * 5 / 11
		}
	}
}
