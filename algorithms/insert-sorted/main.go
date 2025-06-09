package main

import (
	"fmt"
)

func main() {
	var scores = []int{80, 70, 60, 50, 95}
	fmt.Println("Original:", scores)
	fmt.Println("sorted:  ", insertSorted(scores))
}

func insertSorted(scores []int) []int {
	var length = len(scores)

	for i := 0; i < length; i++ {
		element := scores[i]
		position := i
		for j := position - 1; j >= 0; j-- {
			if element < scores[j] {
				scores[j+1] = scores[j]
				position--
			}
		}
		scores[position] = element
	}

	return scores
}
