package main

import (
	"fmt"
	"time"

	"golang.org/x/tour/tree"
)

func sum(s []int, c chan int) {
	sum := 0

	for _, v := range s {
		sum += v
	}

	c <- sum
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("I am boring ...")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// Walk is a demo
func Walk(t *tree.Tree, ch chan int) {
	doWalk(t, ch)
	close(ch)
}

func doWalk(t *tree.Tree, ch chan int) {
	if t != nil {
		doWalk(t.Left, ch)
		ch <- t.Value
		doWalk(t.Right, ch)
	}
}

// Same is a demo
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}

	return true
}

func main9() {
	ch := make(chan int)

	go Walk(tree.New(1), ch)

	for i := range ch {
		fmt.Println(i)
	}

	fmt.Println(Same(tree.New(1), tree.New(1))) // true
	fmt.Println(Same(tree.New(1), tree.New(2))) // false
	fmt.Println(Same(tree.New(3), tree.New(3))) // true
	fmt.Println(Same(tree.New(5), tree.New(4))) // false
	fmt.Println(Same(tree.New(5), tree.New(5))) // true
}
