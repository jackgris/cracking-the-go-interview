package main

import (
	"fmt"
	"math"
)

/*
   https://leetcode.com/problems/divide-two-integers/description/
   Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.

   The integer division should truncate toward zero, which means losing its fractional part. For example, 8.345 would be truncated to 8, and -2.7335 would be truncated to -2.

   Return the quotient after dividing dividend by divisor.

   Note: Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−2^31, 2^31 − 1]. For this problem, if the quotient is strictly greater than 2^31 - 1, then return 2^31 - 1, and if the quotient is strictly less than -2^31, then return -2^31.

   Constraints:

    -2^31 <= dividend, divisor <= 2^31 - 1
    divisor != 0

*/

func main() {
	dividend := 10
	divisor := 3
	output := 3
	// Explanation: 10/3 = 3.33333.. which is truncated to 3.
	fmt.Printf("dividend %d divisor %d output %d result %d\n", dividend, divisor, output, divide(dividend, divisor))

	dividend = 7
	divisor = -3
	output = -2
	// Explanation: 7/-3 = -2.33333.. which is truncated to -2.
	fmt.Printf("dividend %d divisor %d output %d result %d\n", dividend, divisor, output, divide(dividend, divisor))

}

func divide(dividend int, divisor int) int {
	isNegativeResult := (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0)
	dividend = int(math.Abs(float64(dividend)))
	divisor = int(math.Abs(float64(divisor)))

	// related to the constrain: -2^31 <= dividend, divisor <= 2^31 - 1
	upperLimit := int(math.Pow(2, 31) - 1)
	lowerLimit := int(math.Pow(2, 31))
	lowerLimit *= -1

	quotient := 0
	for i := 31; i >= 0; i-- {
		if divisor<<i <= dividend {
			dividend -= divisor << i
			quotient += 1 << i
		}
	}

	if isNegativeResult {
		quotient *= -1
	}

	if quotient > upperLimit {
		return upperLimit
	}

	if quotient < lowerLimit {
		return lowerLimit
	}

	return quotient
}
