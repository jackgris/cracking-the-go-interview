package main

import (
	"fmt"
)

/*
Zigzag Conversion

	The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

	P   A   H   N
	A P L S I I G
	Y   I   R

	And then read line by line: "PAHNAPLSIIGYIR"

	Write the code that will take a string and make this conversion given a number of rows:

	string convert(string s, int numRows);



	Example 1:

	Input: s = "PAYPALISHIRING", numRows = 3
	Output: "PAHNAPLSIIGYIR"

	Example 2:

	Input: s = "PAYPALISHIRING", numRows = 4
	Output: "PINALSIGYAHRPI"

Explanation:

	P     I    N
	A   L S  I G
	Y A   H R
	P     I

	Example 3:

	Input: s = "A", numRows = 1
	Output: "A"

Constraints:

	1 <= s.length <= 1000
	s consists of English letters (lower-case and upper-case), ',' and '.'.
	1 <= numRows <= 1000
*/
func main() {
	s := "PAYPALISHIRING"
	rows := 3
	rightResponse := "PAHNAPLSIIGYIR"
	result := convert(s, rows)
	fmt.Printf("convert receive %s with %d rows and should return %s we get like result %s\n\n", s, rows, rightResponse, result)

	rows = 4
	rightResponse = "PINALSIGYAHRPI"
	result = convert(s, rows)
	fmt.Printf("convert receive %s with %d rows and should return %s we get like result %s\n\n", s, rows, rightResponse, result)

	s = "A"
	rows = 1
	rightResponse = "A"
	result = convert(s, rows)
	fmt.Printf("convert receive %s with %d rows and should return %s we get like result %s\n\n", s, rows, rightResponse, result)

	s = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rows = 5
	rightResponse = "AIQYBHJPRXZCGKOSWDFLNTVEMU"
	result = convert(s, rows)
	fmt.Printf("convert receive %s with %d rows and should return %s we get like result %s\n\n", s, rows, rightResponse, result)

}

func convert(s string, numRows int) string {
	//////////////////////////////////////////////////////////
	// First solution
	//////////////////////////////////////////////////////////
	// if numRows == 1 || len(s) <= numRows {
	// 	return s
	// }

	// res := bytes.Buffer{}
	// // p is the step distance
	// p := numRows*2 - 2

	// // Process the first line
	// for i := 0; i < len(s); i += p {
	// 	res.WriteByte(s[i])
	// }

	// // Process the middle row
	// for r := 1; r <= numRows-2; r++ {
	// 	//Add the first character of line r
	// 	res.WriteByte(s[r])

	// 	for k := p; k-r < len(s); k += p {
	// 		res.WriteByte(s[k-r])
	// 		if k+r < len(s) {
	// 			res.WriteByte(s[k+r])
	// 		}
	// 	}
	// }

	// // Process the last line
	// for i := numRows - 1; i < len(s); i += p {
	// 	res.WriteByte(s[i])
	// }

	// return res.String()

	//////////////////////////////////////////////////////////
	// Second solution
	//////////////////////////////////////////////////////////
	if numRows <= 1 {
		return s
	}

	// character distance between columns
	colStep := numRows + numRows - 2

	// result
	ss := make([]uint8, len(s))
	// fill result with dots to make debugging easier
	//for i := 0; i < len(ss); i++ {
	//	ss[i] = '.'
	//}

	// diagStep is distance from column to diagonal value.
	// Reduces by two for each row
	diagStep := colStep - 2

	// i = position to write to in ss
	i := 0
	for row := 0; row < numRows; row = row + 1 {
		// does this row have diagonal cells?
		diag := row > 0 && row < numRows-1
		for j := row; j < len(ss); j += colStep {
			ss[i] = s[j] // column value
			i += 1
			if diag && j+diagStep < len(s) {
				ss[i] = s[j+diagStep] // diagonal value
				i += 1
			}
		}
		if diag {
			diagStep -= 2
		}
	}

	return string(ss)
}
