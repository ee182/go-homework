package sortalgorithm

import (
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().Unix()))

// QuickSort sorts the array of integers by an randomized algorithm
func QuickSort(xs []int) {
	if len(xs) <= 1 {
		return
	}

	pivotIdx := r.Intn(len(xs))

	p := partition(xs, pivotIdx)
	QuickSort(xs[0:p])
	QuickSort(xs[p+1 : len(xs)])
}

func partition(xs []int, pivotIdx int) int {
	pivot := xs[pivotIdx]
	i := 0

	swap(xs, pivotIdx, 0)

	for j := 1; j < len(xs); j++ {
		if xs[j] < pivot {
			swap(xs, i+1, j)
			i++
		}
	}

	swap(xs, i, 0)

	return i
}

func swap(xs []int, i int, j int) {
	xs[i], xs[j] = xs[j], xs[i]
}
