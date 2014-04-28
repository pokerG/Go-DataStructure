package GoSort

var (
	n, sorted int
	ch        chan []int
	done      = make(chan bool)
)

func swap(a, b *int) {
	c := *a
	(*a) = (*b)
	(*b) = c
}

func max(args ...int) int {
	max := args[0]
	for _, v := range args {
		if v > max {
			max = v
		}
	}
	return max
}
