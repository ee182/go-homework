package sortalgorithm

// QuickSort sorts the array of integers by an randomized algorithm
func QuickSort(xs []int) {
	if len(xs) <= 1 {
		return
	}

	pivotIdx := 0

	p := partition(xs, pivotIdx)
	QuickSort(xs[0:p])
	QuickSort(xs[p+1 : len(xs)])
}

func partition(xs []int, pivotIdx int) int {
	pivot := xs[pivotIdx]
	i := -1

	swap(xs, pivotIdx, len(xs)-1)

	for j := 0; j < len(xs)-1; j++ {
		if xs[j] < pivot {
			swap(xs, i+1, j)
			i++
		}
	}

	swap(xs, i+1, len(xs)-1)

	return i + 1
}

func swap(xs []int, i int, j int) {
	xs[i], xs[j] = xs[j], xs[i]
}
