package sortalgorithm

// MergeSort sorts the array of integers
func MergeSort(xs []int) []int {
	if len(xs) == 1 {
		return xs
	}

	left := MergeSort(xs[0 : len(xs)/2])
	right := MergeSort(xs[len(xs)/2 : len(xs)])

	return Merge(left, right)
}

// Merge merges two sorted array into one sorted array
func Merge(xs []int, ys []int) []int {
	nx := len(xs)
	ny := len(ys)
	n := nx + ny

	ans := make([]int, n, n)
	i, j := 0, 0

	for k := 0; k < n; k++ {
		if i >= nx {
			copy(ans[k:n], ys[j:ny])
			break
		}

		if j >= ny {
			copy(ans[k:n], xs[i:nx])
			break
		}

		if xs[i] < ys[j] {
			ans[k] = xs[i]
			i++
		} else {
			ans[k] = ys[j]
			j++
		}
	}

	return ans
}
