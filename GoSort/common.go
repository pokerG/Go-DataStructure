package GoSort

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
