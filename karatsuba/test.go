package karatsuba

import (
	"fmt"
)

// Test karatsuba multiplication
func Test() {
	a := []int{5, 7, 8, 9, 3, 4, 5, 3, 7, 1}
	b := []int{1, 3, 4, 5, 5, 6, 1, 4, 1, 3}

	fmt.Print(Multiplication(a, b))
	fmt.Print("  ")
	fmt.Println((ToNumbers(a) * ToNumbers(b)))
	fmt.Println(ToNumbers(Multiplication(a, b)) == (ToNumbers(a) * ToNumbers(b)))
}
