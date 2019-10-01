package sortalgorithm

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Test2 is used to test sortalgorithm.QuickSort()
func Test2() {
	xs := []int{3, -9, -10, 8, 2, 22, 3, 22, -100, -100, 100, 12, 12, 13, 14, 15, 15, 14, 13, 12, -1, -2, -3, -4, -5, -1, -5, -4, -3, -2, 12, 3, 4, 5, 6, 7, 5, 1, 44, 4, 44, 7, 6, -9, 10, -10, -15, 16, 17, -1, -15, -1}

	QuickSort(xs)
	fmt.Println(xs)
}

// Test is used to test sortalgorithm.SortAndCount()
func Test() {
	f, err := os.Open("LargeIntegerArray.txt")
	if err != nil {
		fmt.Println("error opening file= ", err)
		os.Exit(1)
	}

	r := bufio.NewReader(f)
	line, err := r.ReadString('\n')

	numbers := make([]int, 0, 0)

	for err == nil {
		i, _ := strconv.Atoi(strings.TrimSuffix(line, "\r\n"))

		numbers = append(numbers, i)

		line, err = r.ReadString('\n')
	}

	fmt.Println(SortAndCount(numbers))
}
