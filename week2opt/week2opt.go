package week2opt

// ValueEqIndex checks whether or not there is an index i such that A[i] = i
func ValueEqIndex(a []int, idx []int) (bool, int) {
	if len(a) == 0 {
		return false, -1
	}

	if len(a) == 1 {
		return a[0] == idx[0], idx[0]
	}

	var middleI = len(a) / 2

	if a[middleI] > idx[middleI] {
		return ValueEqIndex(a[0:middleI], idx[0:middleI])
	}

	if a[middleI] < idx[middleI] {
		return ValueEqIndex(a[middleI+1:len(a)], idx[middleI+1:len(a)])
	}

	return true, idx[middleI]
}
