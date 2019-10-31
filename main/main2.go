package main

import (
	"fmt"
	"math/cmplx"
)

// Those are my factored var
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// Pi is for test
const (
	Pi = 3.14
)

// Those are for test, too
const (
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

// Big is for creatig a huge number by shifting a 1 bit left 100 places.
// In other words, the binary number that is 1 followed by 100 zeroes.
const Big = 1 << 100

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main2() {
	// rand.Seed(time.Now().Unix())
	// fmt.Println("My favorite number is", rand.Intn(10))
	// sortalgorithm.Test2()
	// myfolder01.TestPkg()
	// mypackage02.TestAaaaa()

	// fmt.Println(add(42.5, 13.8))

	// a, b := swap("hello", "world")
	// fmt.Println(a, b)

	// fmt.Println(split(17))

	// var i, j int = 1, 2
	// k := 3
	// c, python, java := true, false, "no!"

	// fmt.Println(i, j, k, c, python, java)

	// fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	// fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	// fmt.Printf("Type: %T Value: %v\n", z, z)

	// var i int
	// var f float64
	// var b bool
	// var s string
	// fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// var x, y int = 3, 4
	// f := math.Sqrt(float64(x*x + y*y))
	// z := uint(f)
	// fmt.Println(x, y, z)

	// v := 0.867 + 0.5i // change me!
	// fmt.Printf("v is of type %T\n", v)

	// const World = "世界"
	// fmt.Println("Hello", World)
	// fmt.Println("Happy", Pi, "Day")

	// const Truth = true
	// fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	// fmt.Println(needInt(Big))
}

func add(x, y float32) float32 {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	return x, y
}
