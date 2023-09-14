package main

import "fmt"

/*
Longest Substring Without Repeating Characters

Given a string s, find the length of the longest
substring
without repeating characters.



	Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.



	Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.
*/

func main() {
	// Input: s = "abcabcbb"
	// Output: 3
	fmt.Printf("RESULT: %d EXPECTED: %d\n\n", lengthOfLongestSubstring("abcabcbb"), 3)
	// Input: s = "bbbbb"
	// Output: 1
	fmt.Printf("RESULT: %d EXPECTED: %d\n\n", lengthOfLongestSubstring("bbbbb"), 1)
	// Input: s = "pwwkew"
	// Output: 3
	fmt.Printf("RESULT: %d EXPECTED: %d\n\n", lengthOfLongestSubstring("pwwkew"), 3)
	// Input: s = "jbpnbwwd"
	// Output: 4
	fmt.Printf("RESULT: %d EXPECTED: %d\n\n", lengthOfLongestSubstring("jbpnbwwd"), 4)
}

func lengthOfLongestSubstring(s string) int {
	// Not optimized but easier to reason - need more time and memory
	// longest := 0
	// count := 0
	// check := map[rune]int{}
	// for p := range s {
	// 	substring := s[p:]
	// 	for _, c := range substring {
	// 		_, ok := check[c]
	// 		if ok {
	// 			count = 0
	// 			check = map[rune]int{}
	// 			break
	// 		} else {
	// 			count++
	// 			check[c] = 1
	// 		}
	// 		if count > longest {
	// 			longest = count
	// 		}
	// 	}
	// }
	// return longest

	////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// This alternative is harder to see because works with integers without converting
	// to string so you are working with integers and thinking about the number of characters that ASCII has.
	////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// location[s[i]] == j means:
	//The i-th string in s last appeared at the j position of s, so there is no s[i] in s[j+1:i]
	// location[s[i]] == -1 means: s[i] appears for the first time in s
	location := [256]int{} // It is only 256 long because it is assumed that the input string has only ASCII characters.
	for i := range location {
		location[i] = -1 // First set all characters that have not been seen before
	}

	maxLen, left := 0, 0

	for i := 0; i < len(s); i++ {
		// Indicates that s[i] has been repeated in s[left:i+1]
		// And the last location of s[i] is location[s[i]]
		if location[s[i]] >= left {
			left = location[s[i]] + 1 // Remove the s[i] character and the part before it in s[left:i+1]
		} else if i+1-left > maxLen {
			maxLen = i + 1 - left
		}
		location[s[i]] = i
	}

	return maxLen
}
