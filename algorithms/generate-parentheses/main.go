package main

import "fmt"

/*
   Generate Parentheses

   Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.


   Example 1:

   Input: n = 3
   Output: ["((()))","(()())","(())()","()(())","()()()"]

   Example 2:

   Input: n = 1
   Output: ["()"]

Constraints:

    1 <= n <= 8
*/

func main() {
	input := 3
	output := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	fmt.Printf("Input: %d Ouput: %s Get: %s\n", input, output, generateParenthesis(input))
	input = 1
	output = []string{"()"}
	fmt.Printf("Input: %d Ouput: %s Get: %s\n", input, output, generateParenthesis(input))
}

func generateParenthesis(n int) []string {
	return backtrack([]string{}, "", 0, 0, n)
}

func backtrack(combinations []string, str string, open int, close int, n int) []string {
	if len(str) == 2*n {
		return append(combinations, str)
	}
	if open < n {
		combinations = backtrack(combinations, str+"(", open+1, close, n)
	}

	if close < open {
		combinations = backtrack(combinations, str+")", open, close+1, n)
	}
	return combinations
}
