package GoSort

import (
	// "fmt"
	"runtime"
)

func QuickSortP(A []int) {
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
			go pqsort()
		}

	}

}
func pqsort() {
	var (
		y, x, tmp int
		slc       []int
	)

	slc = <-ch
	tmp = slc[0]
	y = len(slc)
	x = 0
	for y > x {
		y--
		for (y > x) && (slc[y] >= tmp) {
			y--
		}
		slc[x] = slc[y]
		if y <= x {
			break
		}
		x++
		for (y > x) && (slc[x] <= tmp) {
			x++
		}
		slc[y] = slc[x]
	}
	slc[x] = tmp

	if x > 0 {
		ch <- slc[0:x]
	} else {
		sorted++
	}
	if x < len(slc)-1 {
		ch <- slc[x+1 : len(slc)]
	} else {
		sorted++
	}

	if sorted >= n {
		done <- true
	}

}
