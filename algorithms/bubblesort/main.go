package main

import "fmt"

// Running bubble sort solution
func main() {
	integers := []int{31, 13, 32, 12, 4, 18, 16, 7, 2, 3, 0, 10}
	fmt.Println("Bubble Sorter")
	fmt.Println("Original: ", integers)
	fmt.Println("Result:   ", bubbleSorter(integers))
}

// bubble Sorter method
func bubbleSorter(integers []int) []int {
	isSwapped := true
	for isSwapped {
		isSwapped = false
		for i := 1; i < len(integers); i++ {
			if integers[i-1] > integers[i] {

				temp := integers[i]
				integers[i] = integers[i-1]
				integers[i-1] = temp
				isSwapped = true
			}
		}
	}
	return integers
}
