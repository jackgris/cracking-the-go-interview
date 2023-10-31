package main

import (
	"fmt"
)

/*
Container With Most Water
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.

Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

Example 2:

Input: height = [1,1]
Output: 1

Constraints:

	n == height.length
	2 <= n <= 10^5
	0 <= height[i] <= 10^4
*/
func main() {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	output := 49
	fmt.Printf("With input %v should get %d and we get %d\n", input, output, maxArea(input))
	input = []int{1, 1}
	output = 1
	fmt.Printf("With input %v should get %d and we get %d\n", input, output, maxArea(input))
}

func maxArea(height []int) int {
	n := len(height)
	area, left := 0, 0
	right := n - 1

	for left < right {
		// To calculate the width of the container, find the difference between the right position and the left position in the array.
		width := right - left
		// To calculate the new area, use the formula: 'area = minimun(height1, height2) * width.
		newArea := min(height[left], height[right]) * width
		// Only save the new area if it is larger than the previously saved area.
		area = max(area, newArea)
		// Change the position of the smallest height only; we aim to maintain the higher position because it generates a larger container.
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return area
}

// Time Complexity: O(n)

// Space Complexity: O(1)
