package week2opt

import "github.com/cznic/mathutil"

// Find2ndLargest finds the second-largest number in the array,
// and that uses at most n + lg n - 2 comparisons.
// 1.
func Find2ndLargest(a []int) int {
	cnt := len(a) - 1
	larges := make([]int, 0, cnt)
	smalls := make([]int, 0, cnt)
	backs := make([]int, 0, cnt)

	depth := mathutil.Log2Uint32(uint32(len(a)))

	splitTarget := a
	prevBase := 0
	for j, base, length := 0, 0, len(a)/2; j < depth; j, base, length = j+1, base+length, length/2 {
		largePiece, smallPiece, backPiece := splitLargeSmall(splitTarget, prevBase)
		larges = append(larges, largePiece...)
		smalls = append(smalls, smallPiece...)
		backs = append(backs, backPiece...)

		splitTarget = larges[base : base+length]
		prevBase = base
	}

	seconds := make([]int, depth)
	seconds[0] = smalls[cnt-1]
	backIdx := backs[cnt-1]
	for i := 1; i < depth; i++ {
		seconds[i] = smalls[backIdx]
		backIdx = backs[backIdx]
	}

	return findMax(seconds)
}

func findMax(a []int) int {
	max := a[0]

	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}

	return max
}

func splitLargeSmall(a []int, prevBase int) ([]int, []int, []int) {
	larges := make([]int, len(a)/2)
	smalls := make([]int, len(a)/2)
	backs := make([]int, len(a)/2)

	for i, k := 0, 0; i < len(a); i, k = i+2, k+1 {
		if a[i] > a[i+1] {
			larges[k] = a[i]
			smalls[k] = a[i+1]
			backs[k] = i + prevBase
		} else {
			larges[k] = a[i+1]
			smalls[k] = a[i]
			backs[k] = i + 1 + prevBase
		}
	}

	return larges, smalls, backs
}

// ValueEqIndex checks whether or not there is an index i such that A[i] = i
// 3.
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
