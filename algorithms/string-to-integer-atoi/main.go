package main

import (
	"fmt"
	"math"
)

/*
String to Integer (atoi)

Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer (similar to C/C++'s atoi function).

The algorithm for myAtoi(string s) is as follows:

    Read in and ignore any leading whitespace.
    Check if the next character (if not already at the end of the string) is '-' or '+'. Read this character in if it is either. This determines if the final result is negative or positive respectively. Assume the result is positive if neither is present.
    Read in next the characters until the next non-digit character or the end of the input is reached. The rest of the string is ignored.
    Convert these digits into an integer (i.e. "123" -> 123, "0032" -> 32). If no digits were read, then the integer is 0. Change the sign as necessary (from step 2).
    If the integer is out of the 32-bit signed integer range [-231, 231 - 1], then clamp the integer so that it remains in the range. Specifically, integers less than -231 should be clamped to -231, and integers greater than 231 - 1 should be clamped to 231 - 1.
    Return the integer as the final result.

Note:

    Only the space character ' ' is considered a whitespace character.
    Do not ignore any characters other than the leading whitespace or the rest of the string after the digits.



Example 1:

Input: s = "42"
Output: 42
Explanation: The underlined characters are what is read in, the caret is the current reader position.
Step 1: "42" (no characters read because there is no leading whitespace)
         ^
Step 2: "42" (no characters read because there is neither a '-' nor '+')
         ^
Step 3: "42" ("42" is read in)
           ^
The parsed integer is 42.
Since 42 is in the range [-231, 231 - 1], the final result is 42.

Example 2:

Input: s = "   -42"
Output: -42
Explanation:
Step 1: "   -42" (leading whitespace is read and ignored)
            ^
Step 2: "   -42" ('-' is read, so the result should be negative)
             ^
Step 3: "   -42" ("42" is read in)
               ^
The parsed integer is -42.
Since -42 is in the range [-231, 231 - 1], the final result is -42.

Example 3:

Input: s = "4193 with words"
Output: 4193
Explanation:
Step 1: "4193 with words" (no characters read because there is no leading whitespace)
         ^
Step 2: "4193 with words" (no characters read because there is neither a '-' nor '+')
         ^
Step 3: "4193 with words" ("4193" is read in; reading stops because the next character is a non-digit)
             ^
The parsed integer is 4193.
Since 4193 is in the range [-231, 231 - 1], the final result is 4193.



Constraints:

    0 <= s.length <= 200
    s consists of English letters (lower-case and upper-case), digits (0-9), ' ', '+', '-', and '.'.
*/

func main() {
	s := "42"
	expect := 42
	fmt.Printf("Input %s - Expect %d Result %d\n\n", s, expect, myAtoi(s))

	s = "   -42"
	expect = -42
	fmt.Printf("Input %s - Expect %d Result %d\n\n", s, expect, myAtoi(s))

	s = "4193 with words"
	expect = 4193
	fmt.Printf("Input %s - Expect %d Result %d\n\n", s, expect, myAtoi(s))
}

var digits = map[byte]int{
	0x30: 0,
	0x31: 1,
	0x32: 2,
	0x33: 3,
	0x34: 4,
	0x35: 5,
	0x36: 6,
	0x37: 7,
	0x38: 8,
	0x39: 9,
}

func myAtoi(str string) int {
	res, sign, len, idx := 0, 1, len(str), 0

	// Skip leading spaces
	for idx < len && (str[idx] == ' ' || str[idx] == '\t') {
		idx++
	}

	if idx == len {
		return 0
	}

	// +/- Sign
	if str[idx] == '+' {
		sign = 1
		idx++
	} else if str[idx] == '-' {
		sign = -1
		idx++
	}

	// Digits: 0x30 = '0', 0x31 = '1', ... 0x39 = '9'
	for idx < len && str[idx] >= 0x30 && str[idx] <= 0x39 {
		res = res*10 + digits[str[idx]]
		if sign*res > math.MaxInt32 {
			return math.MaxInt32
		}

		if sign*res < math.MinInt32 {
			return math.MinInt32
		}

		idx++
	}

	return res * sign
}

///////////////////////////////////////////////////////////////////////////////////////////////
// Other solutions
///////////////////////////////////////////////////////////////////////////////////////////////

// func myAtoi(s string) int {
// 	// Remove any additional spaces before and after the given string
// 	s = strings.Trim(s, " ")
// 	n := len(s)
// 	// If string is empty return 0
// 	if n == 0 {
// 		return 0
// 	}
// 	// String index from where the processing should start
// 	var start int
// 	// Handling numbers with +/- sign
// 	signMultiplier := 1
// 	if s[0] == '-' {
// 		// Handling for negative numbers
// 		signMultiplier = -1
// 		start = 1
// 	} else if s[0] == '+' {
// 		// Handling for signed positive number
// 		start = 1
// 	}
// 	// ASCII of '0' = 48
// 	// s[i] - '0' gives the actual number as an integer
// 	var res int
// 	for i := start; i < len(s); i++ {
// 		// Handling for non numeric values
// 		if !(s[i] >= '0' && s[i] <= '9') {
// 			return res * signMultiplier
// 		}
// 		// Calculate the result for current iteration
// 		res = res*10 + int(s[i]-'0')

// 		// Check if result exceeds MinInt32 or MaxInt32
// 		if res*signMultiplier <= math.MinInt32 {
// 			return math.MinInt32
// 		} else if res*signMultiplier >= math.MaxInt32 {
// 			return math.MaxInt32
// 		}
// 	}
// 	return res * signMultiplier
// }

///////////////////////////////////////////////////////////////////////////////////////////////

// func myAtoi(s string) int {
// 	if len(s) == 0 {
// 		return 0
// 	}
// 	sign := 1 // -1 negative
// 	index := 0
// 	val  := 0
// 	for len(s) > index && s[index] == ' ' {
// 		index++;
// 	}
// 	if index >= len(s) {
// 		return 0
// 	}
// 	if v := s[index]; v == '-' || v == '+' {
// 		if v == '-' {
// 			sign = -1
// 		}
// 		index++
// 	}
// 	for {
// 		if index == len(s) {
// 			break
// 		}
// 		if isNumericRune(s[index]) == false {
// 			return sign * val
// 		}

// For those who is wondering how int(int(s[index])-'0') even work, short explanation:
// "subtracting the value of rune '0' from any rune '0' through '9' will give you an integer 0 through 9".
// Resulting type after subtraction int(s[index])-'0' will be int32 (base type of runes)
// 		val = val * 10 + int(s[index]) - '0';

// 		if val > 2147483647 || val< -2147483648 {
// 			if sign == 1 {
// 				return 2147483647
// 			}
// 			return -2147483648
// 		}
// 		index++
// 	}
// 	return sign * val
// }
// func isNumericRune(x byte) bool {
// 	return x >= '0' && x <= '9'
// }
