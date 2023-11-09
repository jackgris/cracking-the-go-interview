package main

import (
	"fmt"
)

/*
  Letter Combinations of a Phone Number
  Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

   A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

   2 = a b c
   3 = d e f
   4 = g h i
   5 = j k l
   6 = m n o
   7 = p q r s
   8 = t u v
   9 = w x y z

   Example 1:

   Input: digits = "23"
   Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

   Example 2:

   Input: digits = ""
   Output: []

   Example 3:

   Input: digits = "2"
   Output: ["a","b","c"]


   Constraints:

    0 <= digits.length <= 4
    digits[i] is a digit in the range ['2', '9'].
*/

func main() {

	input := "23"
	output := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	fmt.Printf("Input %s output should be %s and the result is: %s\n", input, output, letterCombinations(input))

	input = ""
	output = []string{}
	fmt.Printf("Input %s output should be %s and the result is: %s\n", input, output, letterCombinations(input))

	input = "2"
	output = []string{"a", "b", "c"}
	fmt.Printf("Input %s output should be %s and the result is: %s\n", input, output, letterCombinations(input))
}

/*
	  Recursive Solution:

		func letterCombinations(digits string) []string {
			if digits == "" {
				return []string{}
			}

			phoneMap := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
			var output []string

			var backtrack func(combination string, nextDigits string)
			backtrack = func(combination string, nextDigits string) {
				if nextDigits == "" {
					output = append(output, combination)
				} else {
					letters := phoneMap[nextDigits[0]-'2']
					for _, letter := range letters {
						backtrack(combination+string(letter), nextDigits[1:])
					}
				}
			}

			backtrack("", digits)
			return output
		  }
*/

var alpha = map[rune]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	size := 1
	for _, d := range digits {
		size *= len(alpha[d])
	}
	res := make([]string, size)
	for _, d := range digits {
		curr := alpha[d]
		size /= len(curr)
		for i := range res {
			res[i] += string(curr[(i/size)%len(curr)])
		}
	}
	return res
}
