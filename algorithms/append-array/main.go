package main

import "fmt"

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}
	fmt.Println("Original:             ", scores)
	fmt.Println("Array with new value: ", appendValue(scores, 75))
}

func appendValue(scores []int, value int) []int {
	length := len(scores)
	var newScores = make([]int, length+1)
	for i := 0; i < len(scores); i++ {
		newScores[i] = scores[i]
	}
	newScores[length] = value
	return newScores
}
