package main

import (
	"fmt"
	"math"
)

// Pow is for demo
func Pow(x, n, lim float64) float64 {
	v := math.Pow(x, n)
	if v < lim {
		return v
	}

	fmt.Printf("%g >= %g\n", v, lim)

	return lim
}

// Sqrt using gradient descent to approximate sqrt
func Sqrt(x float64) float64 {
	z := x / 2
	// z1 := x
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z) // Newton's method
		// z1 -= 0.00001 * 2 * 2 * (z1*z1 - x) * z1 // Gradient descent, J(z) = (z^2 - x)^2
		// fmt.Println(z)
	}

	return z
}

func main3() {
	for x := 1.0; x < 257.0; x += 1.0 {
		fmt.Println(Sqrt(x))
	}
}
