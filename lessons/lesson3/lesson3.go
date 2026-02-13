package lesson3

import "fmt"

func Run() {
	// numList := []int{1, 2, 3, 4, 5}
	// fmt.Println(Sum(numList))
	fmt.Println(max(1, 2))
}

func Sum(numList []int) int {
	sum := 0
	for _, num := range numList {
		sum += num
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}