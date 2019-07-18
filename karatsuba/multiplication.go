package karatsuba

import "math"

// Multiplication will give you the multiplication
// of two numbers, by means of a smart recursive algorithm
func Multiplication(a []int, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return []int{0}
	}

	if len(a) == 1 && len(b) == 1 {
		return ToDigits(a[0] * b[0])
	}

	var k = min(len(a)/2, len(b)/2)
	if k == 0 {
		k = 1
	}

	var high = k * 2

	var A = a[0 : len(a)-k]
	var B = a[len(a)-k : len(a)]
	var C = b[0 : len(b)-k]
	var D = b[len(b)-k : len(b)]

	var AC = Multiplication(A, C)
	var BD = Multiplication(B, D)
	var ApBCpD = Multiplication(add(A, B), add(C, D))
	var ADpBC = sub(sub(ApBCpD, AC), BD)

	return add(padZeros(AC, high), add(padZeros(ADpBC, k), BD))
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

// TODO: do add() digit by digit instead of using native arithmetics
func add(a []int, b []int) []int {
	return ToDigits(ToNumbers(a) + ToNumbers(b))
}

// TODO: do sub() digit by digit instead of using native arithmetics
func sub(a []int, b []int) []int {
	return ToDigits(ToNumbers(a) - ToNumbers(b))
}

// ToDigits ... remove please
func ToDigits(i int) []int {
	ans := []int{}
	for i > 0 {
		ans = append([]int{i - i/10*10}, ans...)
		i = i / 10
	}

	return ans
}

// ToNumbers ... remove please
func ToNumbers(i []int) int {
	var ans = 0
	for index, element := range i {
		ans = ans + element*int(math.Pow10(len(i)-index-1))
	}

	return ans
}

func padZeros(a []int, n int) []int {
	for i := 0; i < n; i++ {
		a = append(a, 0)
	}

	return a
}
