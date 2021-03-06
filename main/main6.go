package main

import (
	"fmt"
	"math"
)

// I is a demo
type I interface{}

// Vertex is a demo
type Vertex struct {
	X, Y float64
}

// Abs is a demo
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func ff(args ...string) {
	fmt.Println(len(args))
}

func main6() {
	var i I
	describe(i)

	i = 42
	describe(i)

	i = "hallo.quantum"
	describe(i)

	i = &Vertex{182, 183}
	describe(i)

	args := []string{"a", "b", "a", "b", "a", "b", "a", "b"}

	ff(args...)

}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
