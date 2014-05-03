//This package provides some sort way
//
//Copytright (C) 2014 by pokerG <pokerfacehlg@gmail.com>
package GoSort

import (
	"math"
)

// BucketSort
// O(n log (r) m) running time.
const null int = -1

func BucketSort(A []int) {
	l := len(A)
	if l < 2 {
		return
	}
	maxNum := max(A...)
	loopTimes := getLoopTimes(maxNum)
	for i := 1; i <= loopTimes; i++ {
		radix(A, i)
	}
}

func getLoopTimes(num int) int {
	count := 1
	temp := num / 10
	for temp != 0 {
		count++
		temp = temp / 10
	}
	return count
}

func radix(a []int, m int) {
	var buckets [10][]int
	for i := 0; i < 10; i++ {
		buckets[i] = make([]int, len(a))
		for j := 0; j < len(a); j++ {
			buckets[i][j] = null

		}
	}
	tempNum := int(math.Pow10(m - 1)) // 1, 10 ,100...
	for i := 0; i < len(a); i++ {
		rowIndex := (a[i] / tempNum) % 10
		for j := 0; j < len(a); j++ {
			if buckets[rowIndex][j] == null {
				buckets[rowIndex][j] = a[i]
				break
			}
		}
	}
	k := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < len(a); j++ {
			if buckets[i][j] != null {
				a[k] = buckets[i][j]
				k++
			}
		}
	}
}
