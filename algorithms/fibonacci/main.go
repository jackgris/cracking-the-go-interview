package main

import (
	"fmt"
)

// Running fibonacci solutions
func main() {
	fibo := 9
	seriesResult := []int{}
	for i := 0; i <= fibo; i++ {
		seriesResult = append(seriesResult, series(i))
	}
	fmt.Println("Series:    ", seriesResult)
	recursiveResult := []int{}
	for i := 0; i <= fibo; i++ {
		recursiveResult = append(recursiveResult, recursive(i))
	}
	fmt.Println("Recursive: ", recursiveResult)
}

// series method
func series(n int) int {
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

// recursive method
func recursive(n int) int {
	if n <= 1 {
		return n
	}
	return recursive(n-1) + recursive(n-2)
}
