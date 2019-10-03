package sortalgorithm

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Test3 is used to test the probability of sortalgorithm.QuickSort()
// func Test3() {
// 	trueCnt := 0
// 	totalCnt := 1000000
// 	for i := 0; i < totalCnt; i++ {
// 		numbers := []int{6, 2, 5, 1, 4, 8, 3, 7}
// 		compared36 = false
// 		QuickSort(numbers)

// 		if compared36 == true {
// 			trueCnt++
// 		}
// 	}

// 	fmt.Println(float32(trueCnt) / float32(totalCnt))
// }

// Test2 is used to test sortalgorithm.QuickSort()
func Test2() {
	f, err := os.Open("QuickSortNumbers.txt")
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

	QuickSort(numbers)
	fmt.Println(numbers)
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
