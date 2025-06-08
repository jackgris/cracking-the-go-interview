package main

import "fmt"

func main() {
	scores := []int{60, 80, 95, 50, 70}
	fmt.Println("Original: ", scores)
	fmt.Println("Min value:", min(scores))
}

func min(integers []int) int {
	minIndex := 0
	for i := 1; i < len(integers); i++ {
		if integers[minIndex] > integers[i] {
			minIndex = i
		}
	}

	return integers[minIndex]
}
