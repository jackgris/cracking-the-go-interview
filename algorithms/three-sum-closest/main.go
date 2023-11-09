package main

import (
	"fmt"
	"sort"
)

/*
   3Sum Closest

   Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target.

   Return the sum of the three integers.

   You may assume that each input would have exactly one solution.



   Example 1:

   Input: nums = [-1,2,1,-4], target = 1
   Output: 2
   Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

   Example 2:

   Input: nums = [0,0,0], target = 1
   Output: 0
   Explanation: The sum that is closest to the target is 0. (0 + 0 + 0 = 0).



   Constraints:

    3 <= nums.length <= 500
    -1000 <= nums[i] <= 1000
    -10^4 <= target <= 10^4
*/

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	output := 2
	fmt.Printf("With nums %d and a target of %d, we should obtain %d, and the result is: %d\n", nums, target, output, threeSumClosest(nums, target))
	nums = []int{0, 0, 0}
	target = 1
	output = 0
	fmt.Printf("With nums %d and a target of %d, we should obtain %d, and the result is: %d\n", nums, target, output, threeSumClosest(nums, target))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < n-2; i++ {
		j, k := i+1, n-1
		for j < k {
			tmp := nums[i] + nums[j] + nums[k]
			if tmp == target {
				return target
			}
			if abs(target-res) > abs(target-tmp) {
				res = tmp
			}
			if tmp > target {
				k--
			} else {
				j++
			}
		}
	}
	return res
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
