package main

import (
	"fmt"
)

func main() {
	var scores = []int{30, 40, 50, 70, 85, 90, 100}
	value := 100
	fmt.Printf("list %d value %d index %d\n", scores, value, binarySearch(scores, value))
}

func binarySearch(list []int, value int) int {
	low := 0
	hight := len(list) - 1
	if list[low] == value {
		return low
	}

	if list[hight] == value {
		return hight
	}

	for low < hight {
		middle := (low + hight) / 2
		if value == list[middle] {
			return middle
		} else if value < list[middle] {
			hight = middle - 1
		} else {
			low = middle + 1
		}
	}

	return -1
}
