package main

import (
	"fmt"
	"slices"
)

/*
   Given an array of integers nums and an integer threshold, we will choose a positive integer divisor, divide all the array by it, and sum the division's   result. Find the smallest divisor such that the result mentioned above is less than or equal to threshold.

   Each result of the division is rounded to the nearest integer greater than or equal to that element. (For example: 7/3 = 3 and 10/2 = 5).

   The test cases are generated so that there will be an answer.

   Example 1:

   Input: nums = [1,2,5,9], threshold = 6
   Output: 5
   Explanation: We can get a sum to 17 (1+2+5+9) if the divisor is 1.
   If the divisor is 4 we can get a sum of 7 (1+1+2+3) and if the divisor is 5 the sum will be 5 (1+1+1+2).

   Example 2:

   Input: nums = [44,22,33,11,1], threshold = 5
   Output: 44

   Constraints:

   1 <= nums.length <= 5 * 104
   1 <= nums[i] <= 106
   nums.length <= threshold <= 106
*/

func main() {
	input := []int{1, 2, 5, 9}
	threshold := 6
	output := 5
	fmt.Printf("input %d threshold %d output %d result %d\n", input, threshold, output, smallestDivisor(input, threshold))

	input = []int{44, 22, 33, 11, 1}
	threshold = 5
	output = 44
	fmt.Printf("input %d threshold %d output %d result %d\n", input, threshold, output, smallestDivisor(input, threshold))
}

func smallestDivisor(nums []int, threshold int) int {
	// Get the smaller and bigger possibilities.
	smaller, bigger := 1, slices.Max(nums)
	// We will try until we have not other possibility.
	for smaller <= bigger {
		medium := (smaller + bigger) / 2
		result := sumDividedNumbers(nums, medium)
		// After calculating our result, we know we need to try with a smaller medium if the result is smaller than our threshold.
		// If not, we will need a bigger medium value.
		if result <= threshold {
			bigger = medium - 1
		} else {
			smaller = medium + 1
		}
	}
	return smaller
}

func sumDividedNumbers(nums []int, divisor int) int {
	summary := 0
	// The positive integer divisor divides all the array by it and sums the division's result.
	for _, n := range nums {
		summary += n / divisor
		// Each result of the division is rounded to the nearest integer greater than or equal to that element.
		// For example: 7/3 = 3
		if n%divisor != 0 {
			summary++
		}
	}
	return summary
}
