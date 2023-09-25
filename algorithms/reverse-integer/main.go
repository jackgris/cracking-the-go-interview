package main

import (
	"fmt"
	"math"
)

/*
Reverse Integer

Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).



	Example 1:

Input: x = 123
Output: 321

Example 2:

Input: x = -123
Output: -321

Example 3:

Input: x = 120
Output: 21



Constraints:

-231 <= x <= 231 - 1
*/

func main() {
	x := 123
	expected := 321
	result := reverse(x)
	fmt.Printf("Reverse %d should be %d and we get %d\n\n", x, expected, result)

	x = -123
	expected = -321
	result = reverse(x)
	fmt.Printf("Reverse %d should be %d and we get %d\n\n", x, expected, result)

	x = 120
	expected = 21
	result = reverse(x)
	fmt.Printf("Reverse %d should be %d and we get %d\n", x, expected, result)
}

// func reverse(x int) int {
// 	r := ""
// 	if x < 0 {
// 		r += "-"
// 		x -= (x * 2)
// 	}
// 	c := strconv.Itoa(x)
// 	for i := len(c) - 1; i >= 0; i-- {
// 		r += string(c[i])
// 	}
// 	x, _ = strconv.Atoi(r)
// 	if isOverFlow(x) {
// 		return 0
// 	}
// 	return x
// }

// func isOverFlow(x int) bool {
// 	max := math.Pow(2, 31) - 1
// 	min := math.Pow(-2, 31)
// 	y := float64(x)
// 	if min <= y && y <= max {
// 		return false
// 	}
// 	return true
// }

// func reverse(x int) int {
// 	signMultiplier := 1
// 	if x < 0 {
// 		signMultiplier = -1
// 		x = signMultiplier * x
// 	}
// 	var result int
// 	for x > 0 {
// 		// Add the current digit into result
// 		result = result*10 + x%10
// 		// Check if the result is within MaxInt32 and MinInt32 bounds
// 		if signMultiplier*result >= math.MaxInt32 || signMultiplier*result <= math.MinInt32 {
// 			return 0
// 		}
// 		x = x / 10
// 	}
// 	// Restore to original sign of number (+ or -)
// 	return signMultiplier * result
// }

func reverse(x int) int {
	result := 0
	for x != 0 {
		// Extract last digit from x using the mod operator. Example: currentDigit = x%10 = 369 % 10 = 9.
		currentDidit := x % 10
		// Put currentDigit to its correct base position in result. Example: result = result * 10 + currentDigit = 0 * 10 + 9 = 9.
		result = result*10 + currentDidit
		// Now that we have processed the digit we need to remove it from x. Example: x = x/10 = 369/10 = 36.
		x = x / 10
	}

	if result < math.MinInt32 || result > math.MaxInt32 {
		return 0
	}
	return result
}
