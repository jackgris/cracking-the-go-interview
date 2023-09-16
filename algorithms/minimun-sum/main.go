package main

import (
	"sort"
)

// Given an array of integers, perform some number k of operations.
// Each operation consists of removing an element from the array, dividing it by 2, and inserting the ceiling of that result back into the array.
// Minimize the sum of the elements in the final array

func minSum(num []int32, k int32) (sum int32) {

	sort.Slice(num, func(i, j int) bool {
		return num[i] > num[j]
	})

	for i := 0; i < int(k); i++ {
		sorting(num)
		num[0] = (num[0] + 1) / 2
	}

	for _, v := range num {
		sum += v
	}

	return
}

func sorting(num []int32) {
	for i := 0; i < len(num)-1; i++ {
		if num[i] > num[i+1] {
			break
		}
		num[i], num[i+1] = num[i+1], num[i]
	}
}
