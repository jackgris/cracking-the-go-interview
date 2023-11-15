package main

import (
	"fmt"
	"sort"
)

/*
   4Sum

   Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:

    0 <= a, b, c, d < n
    a, b, c, and d are distinct.
    nums[a] + nums[b] + nums[c] + nums[d] == target

   You may return the answer in any order.

   Example 1:

   Input: nums = [1,0,-1,0,-2,2], target = 0
   Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

   Example 2:

   Input: nums = [2,2,2,2,2], target = 8
   Output: [[2,2,2,2]]

   Constraints:

    1 <= nums.length <= 200
    -10^9 <= nums[i] <= 10^9
    -10^9 <= target <= 10^9
*/

func main() {
	input := []int{1, 0, -1, 0, -2, 2}
	target := 0
	output := [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}
	fmt.Printf("Input %d Target %d should Output %d Result %d\n", input, target, output, fourSum(input, target))

	input = []int{2, 2, 2, 2, 2}
	target = 8
	output = [][]int{{2, 2, 2, 2}}
	fmt.Printf("Input %d Target %d should Output %d Result %d\n", input, target, output, fourSum(input, target))
}

func fourSum(nums []int, target int) [][]int {
	result := [][]int{}
	length := len(nums)
	sort.Ints(nums)
	for i := 0; i < length-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		if nums[i]+nums[length-1]+nums[length-2]+nums[length-2] < target {
			continue
		}

		for j := i + 1; j < length-2; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			if nums[i]+nums[j]+nums[length-1]+nums[length-2] < target {
				continue
			}
			head, end := j+1, length-1
			for head < end {
				tempt := nums[i] + nums[j] + nums[head] + nums[end]
				if tempt == target {
					result = append(result, []int{nums[i], nums[j], nums[head], nums[end]})
					head++
					for head < end && nums[head] == nums[head-1] {
						head++
					}
				} else if tempt > target {
					end--
				} else {
					head++
				}
			}
		}
	}
	return result
}
