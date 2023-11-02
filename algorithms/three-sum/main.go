package main

import (
	"fmt"
	"sort"
)

/*

3Sum

Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

   Notice that the solution set must not contain duplicate triplets.



   Example 1:

   Input: nums = [-1,0,1,2,-1,-4]
   Output: [[-1,-1,2],[-1,0,1]]
   Explanation:
   nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
   nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
   nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
   The distinct triplets are [-1,0,1] and [-1,-1,2].
   Notice that the order of the output and the order of the triplets does not matter.

   Example 2:

   Input: nums = [0,1,1]
   Output: []
   Explanation: The only possible triplet does not sum up to 0.

   Example 3:

   Input: nums = [0,0,0]
   Output: [[0,0,0]]
   Explanation: The only possible triplet sums up to 0.



Constraints:

    3 <= nums.length <= 3000
    -10^5 <= nums[i] <= 10^5
*/

func main() {
	input := []int{-1, 0, 1, 2, -1, -4}
	output := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	fmt.Printf("Input %v Output should be %v and we get %v\n", input, output, threeSum(input))

	input = []int{0, 1, 1}
	output = [][]int{}
	fmt.Printf("Input %v Output should be %v and we get %v\n", input, output, threeSum(input))

	input = []int{0, 0, 0}
	output = [][]int{{0, 0, 0}}
	fmt.Printf("Input %v Output should be %v and we get %v\n", input, output, threeSum(input))
}

/*
	func threeSum(nums []int) [][]int {
		var results [][]int
		sort.Ints(nums)
		for i := 0; i < len(nums)-2; i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue //To prevent the repeat
			}
			target, left, right := -nums[i], i+1, len(nums)-1
			for left < right {
				sum := nums[left] + nums[right]
				if sum == target {
					results = append(results, []int{nums[i], nums[left], nums[right]})
					left++
					right--
					for left < right && nums[left] == nums[left-1] {
						left++
					}
					for left < right && nums[right] == nums[right+1] {
						right--
					}
				} else if sum > target {
					right--
				} else if sum < target {
					left++
				}
			}
		}
		return results
	}
*/
func threeSum(nums []int) [][]int {
	n := len(nums)

	// Sort the given array
	sort.Ints(nums)

	var result [][]int
	for num1Idx := 0; num1Idx < n-2; num1Idx++ {
		// Skip all duplicates from left
		// num1Idx>0 ensures this check is made only from 2nd element onwards
		if num1Idx > 0 && nums[num1Idx] == nums[num1Idx-1] {
			continue
		}

		num2Idx := num1Idx + 1
		num3Idx := n - 1
		for num2Idx < num3Idx {
			sum := nums[num2Idx] + nums[num3Idx] + nums[num1Idx]
			if sum == 0 {
				// Add triplet to result
				result = append(result, []int{nums[num1Idx], nums[num2Idx], nums[num3Idx]})

				num3Idx--

				// Skip all duplicates from right
				for num2Idx < num3Idx && nums[num3Idx] == nums[num3Idx+1] {
					num3Idx--
				}
			} else if sum > 0 {
				// Decrement num3Idx to reduce sum value
				num3Idx--
			} else {
				// Increment num2Idx to increase sum value
				num2Idx++
			}
		}
	}
	return result
}
