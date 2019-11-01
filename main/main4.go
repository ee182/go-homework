package main

import (
	"fmt"
)

func f() int {
	fmt.Println("i AM F")

	return 2
}

func old() {
	// fmt.Println("When's Saturday?")
	// today := time.Now().Weekday()
	// switch time.Saturday {
	// case today + 0:
	// 	fmt.Println("Today.")
	// case today + 1:
	// 	fmt.Println("Tomorrow.")
	// case today + 2:
	// 	fmt.Println("In two days.")
	// default:
	// 	fmt.Println("Too far away.")
	// }

	// t := time.Now()
	// switch {
	// case t.Hour() < 12:
	// 	fmt.Println("Good morning!")
	// case t.Hour() < 17:
	// 	fmt.Println("Good afternoon.")
	// default:
	// 	fmt.Println("Good evening.")
	// }

	i := 5
	defer fmt.Println(i)
	defer fmt.Print("world, i am ")

	i = 10

	fmt.Println("hello")
}

// // Vertex is for testing
// type Vertex struct {
// 	X, Y int
// }

// var (
// 	v1 = Vertex{1, 2}   // has type Vertex
// 	v2 = Vertex{Y: 182} // X:0 is implicit
// 	v3 = Vertex{}       // X:0 and Y:0
// 	p  = &Vertex{1, 2}  // has type *Vertex
// )

func main4() {
	// var a [6]string = [6]string{"i", "am", "robot", "you", "know", "?"}
	// a[0] = "Hello"
	// a[1] = "World"
	// fmt.Println(a[0], a[1])
	// fmt.Println(a)
	// fmt.Println(a[2:6])

	// primes := [6]int{2, 3, 5, 7, 11, 13}
	// fmt.Println(primes)

	// s := primes[1:4]
	// fmt.Println(s)

	// q := []int{2, 3, 5, 7, 11, 13}
	// fmt.Println(q)

	// r := []bool{true, false, true, true, false, true}
	// fmt.Println(r)

	// s := []struct {
	// 	i int
	// 	b bool
	// }{
	// 	{2, true},
	// 	{3, false},
	// 	{5, true},
	// 	{7, true},
	// 	{11, false},
	// 	{13, true},
	// }
	// fmt.Println(s)

	// s := []int{2, 3, 5, 7, 11, 13}

	// s = s[1:4]
	// fmt.Println(s)

	// s = s[:2]
	// fmt.Println(s)

	// s = s[1:]
	// fmt.Println(s)

	// s := []int{2, 3, 5, 7, 11, 13}
	// printSlice(s)

	// // Slice the slice to give it zero length.
	// s = s[:0]
	// printSlice(s)

	// // Extend its length.
	// s = s[:4]
	// printSlice(s)

	// // Drop its first two values.
	// s = s[2:]
	// printSlice(s)

	// var s []int
	// fmt.Println(s, len(s), cap(s))
	// if s == nil {
	// 	fmt.Println("nil!")
	// }

	// Create a tic-tac-toe board.
	// board := [][]string{
	// 	[]string{"_", "_", "_"},
	// 	[]string{"_", "_", "_"},
	// 	[]string{"_", "_", "_"},
	// }

	// // The players take turns.
	// board[0][0] = "X"
	// board[2][2] = "O"
	// board[1][2] = "X"
	// board[1][0] = "O"
	// board[0][2] = "X"

	// for i := 0; i < len(board); i++ {
	// 	fmt.Printf("%s\n", strings.Join(board[i], " "))
	// }

	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)

	appendMe := []int{5, 6, 7, 8, 9, 10}

	s = append(s, appendMe...)

	printSlice(s)
}

// func printSlice(s string, x []int) {
// 	fmt.Printf("%s len=%d cap=%d %v\n",
// 		s, len(x), cap(x), x)
// }

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
