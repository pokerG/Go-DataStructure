package GoSort

import (
	"runtime"
)

func MergeSortP(A []int) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	n = len(A)
	ch = make(chan []int, n)
	test := make(chan bool)
	close(test)
	ch <- A
f:
	for {
		select {
		case <-done:
			// fmt.Println("@@@")
			break f
		case <-test:
			// fmt.Println("!!!")
			go msp()
		}

	}

}

func msp() {
	var slc []int
	slc = <-ch
	l := len(slc)
	if l < 2 {
		return
	}
	if l < 7 {
		InsertSort(slc)
		return
	}
	middle := l / 2
	left := slc[:middle]
	right := slc[middle:]
	ch <- left
	ch <- right
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
	copy(slc, result)
	if len(slc) == n {
		done <- true
	}
}
