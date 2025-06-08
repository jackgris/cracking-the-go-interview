package main

import "fmt"

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}
	fmt.Println("Original:           ", scores)
	fmt.Println("One element deleted:", deleteElement(scores, 2))
}

func deleteElement(scores []int, position int) []int {
	if position < 0 || position >= len(scores) {
		// Return the original slice or handle error as needed
		return scores
	}
	var newArray = make([]int, len(scores)-1)

	for i := 0; i < len(scores)-1; i++ {
		if i < position {
			newArray[i] = scores[i]
		} else {
			newArray[i] = scores[i+1]
		}
	}

	// Another option
	// copy(newArray, scores[:position])
	// copy(newArray[position:], scores[position+1:])

	// An other option
	// return append(scores[:position], scores[position+1:]...)

	return newArray
}
