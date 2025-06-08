package main

import "fmt"

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}
	fmt.Println("Original:  ", scores)
	fmt.Println("New array: ", insertOne(scores, 75, 3))
}

func insertOne(scores []int, value, position int) []int {
	var newArray = make([]int, len(scores)+1)
	// It depends on whether we want to insert the value at the index or at the position in the array.
	// In other words, whether we start counting from 0 or 1.
	if position > 0 {
		position -= 1
	}
	for i := 0; i < len(scores); i++ {
		if i < position {
			newArray[i] = scores[i]
		} else {
			newArray[i+1] = scores[i]
		}
	}

	newArray[position] = value

	return newArray
}
