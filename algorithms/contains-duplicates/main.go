package main

/*
   -Level: medium
   -Type: algorithm
   -Description: We want to obtain a complexity of O(n) with a memory consumption of O(1). Complete the function that receives a list with n letters and validates if there are duplicates.
   -Points to consider:
 A lowercase letter is equal to a capital letter (Ex. 'A' == 'a')
   -Function: func containsDuplicates(letters []rune) bool

   Examples:
   [a, A, b] = true
   [a, b, z]   = false

   This will help:
   https://www.ardanlabs.com/blog/2021/04/using-bitmasks-in-go.html
*/

import (
	"fmt"
	"unicode"
)

func main() {
	letters := []rune{'a', 'A', 'b'}
	fmt.Println(containsDuplicates(letters)) // Output: true
	letters = []rune{'a', 'b', 'z'}
	fmt.Println(containsDuplicates(letters)) // Output: false
	letters = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u'}
	fmt.Println(containsDuplicates(letters)) // Output: false
	letters = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'a'}
	fmt.Println(containsDuplicates(letters)) // Output: true
}

func containsDuplicates(letters []rune) bool {
	// Create a bitmask to store the seen characters
	seen := 0

	for _, letter := range letters {
		// We make sure that our rune is lowercase.
		letter := unicode.ToLower(letter)
		// Get the bit position for the current character
		bitPos := 1 << uint(letter-'a') // Because all letters are lowercase
		// Check if the bit is already set
		if (seen & bitPos) > 0 {
			return true
		}

		// Set the bit for the current character
		seen |= bitPos
	}

	return false
}
