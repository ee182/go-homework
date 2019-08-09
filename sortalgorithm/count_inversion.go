package sortalgorithm

// SortAndCount counts the number of inversions and sorts the array of integers
func SortAndCount(xs []int) ([]int, int) {
	if len(xs) == 1 {
		return xs, 0
	}

	left, x := SortAndCount(xs[0 : len(xs)/2])
	right, y := SortAndCount(xs[len(xs)/2 : len(xs)])
	merged, z := MergeAndCount(left, right)

	return merged, x + y + z
}

// MergeAndCount counts the number of split-inversions and merges two sorted array into one sorted array
func MergeAndCount(xs []int, ys []int) ([]int, int) {
	nx := len(xs)
	ny := len(ys)
	n := nx + ny

	splitInv := 0

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

		if xs[i] > ys[j] {
			ans[k] = ys[j]
			splitInv = splitInv + nx - i
			j++
		} else {
			ans[k] = xs[i]
			i++
		}
	}

	return ans, splitInv
}
