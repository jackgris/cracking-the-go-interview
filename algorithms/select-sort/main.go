package main

import "fmt"

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}
	fmt.Println("Array: ", scores)
	fmt.Println("Sorted:", selectSort(scores))
}

func selectSort(scores []int) []int {

	for i := 0; i < len(scores)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(scores); j++ {
			if scores[minIndex] > scores[j] {
				minIndex = j
			}
		}
		if minIndex != i {
			temp := scores[i]
			scores[i] = scores[minIndex]
			scores[minIndex] = temp
		}
	}

	return scores
}
