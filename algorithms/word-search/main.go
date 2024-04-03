package main

import (
	"fmt"
)

/*
 * https://leetcode.com/problems/word-search/

 Given an m x n grid of characters board and a string word, return true if word exists in the grid.

   The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.

   Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
   Output: true

   Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
   Output: true

   Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
   Output: false

    m == board.length
    n = board[i].length
    1 <= m, n <= 6
    1 <= word.length <= 15
    board and word consists of only lowercase and uppercase English letters.

*/

func main() {
	input := [][]string{{"A", "B", "C", "E"}, {"S", "F", "C", "S"}, {"A", "D", "E", "E"}}
	word := "ABCCED"
	output := true
	fmt.Printf("input: %s word: %s output: %v result: %v\n", input, word, output, exist(stringsToBytes(input), word))

	input = [][]string{{"A", "B", "C", "E"}, {"S", "F", "C", "S"}, {"A", "D", "E", "E"}}
	word = "SEE"
	output = true
	fmt.Printf("input: %s word: %s output: %v result: %v\n", input, word, output, exist(stringsToBytes(input), word))

	input = [][]string{{"A", "B", "C", "E"}, {"S", "F", "C", "S"}, {"A", "D", "E", "E"}}
	word = "ABCB"
	output = false
	fmt.Printf("input: %s word: %s output: %v result: %v\n", input, word, output, exist(stringsToBytes(input), word))
}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] && traverse(i, j, 0, board, word) {
				return true
			}
		}
	}

	return false
}

func traverse(x, y, cnt int, board [][]byte, word string) bool {
	if cnt == len(word) {
		return true
	}

	if x < 0 || y < 0 || x == len(board) || y == len(board[0]) || board[x][y] != word[cnt] {
		return false
	}

	state := board[x][y]
	board[x][y] = '_'

	if traverse(x+1, y, cnt+1, board, word) ||
		traverse(x-1, y, cnt+1, board, word) ||
		traverse(x, y+1, cnt+1, board, word) ||
		traverse(x, y-1, cnt+1, board, word) {
		return true
	}
	board[x][y] = state

	return false
}

func stringsToBytes(input [][]string) [][]byte {
	output := [][]byte{}

	for _, i := range input {
		chunk := ""
		for _, j := range i {
			chunk += j
		}
		output = append(output, []byte(chunk))
	}

	return output
}
