package main

import (
	"fmt"
)

/*
Given a string s, return the longest palindromic substring in s.

Example 1:

Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.

Example 2:

Input: s = "cbbd"
Output: "bb"



Constraints:

- 1 <= s.length <= 1000
- s consist of only digits and English letters.
*/

func main() {
	result := LongestPalindrome("babad")
	fmt.Println(result)
	result = LongestPalindrome("cbbd")
	fmt.Println(result)
}

func LongestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	start, maxLength := 0, 1
	for i := 0; i < len(s)-maxLength/2; i++ {
		newStart, newLength := extendPalindrome(s, i, i)
		if newLength > maxLength {
			maxLength = newLength
			start = newStart
		}
		newStart, newLength = extendPalindrome(s, i, i+1)
		if newLength > maxLength {
			maxLength = newLength
			start = newStart
		}
	}
	return s[start : start+maxLength]
}

func extendPalindrome(s string, i, j int) (int, int) {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return i + 1, j - i - 1
}
