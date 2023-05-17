package main

import (
	"fmt"
	"strconv"
)

// Running fibonacci solutions
func main() {
	var i int
	for i = 0; i <= 9; i++ {
		fmt.Print(strconv.Itoa(series(i)) + " ")
	}
	fmt.Println("")
	for i = 0; i <= 9; i++ {
		fmt.Print(strconv.Itoa(fibonacciNumber(i)) + " ")
	}
	fmt.Println("")
}

// series method
func series(n int) int {
	var f []int
	f = make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	var i int
	for i = 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

// fibonacciNumber method
func fibonacciNumber(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciNumber(n-1) + fibonacciNumber(n-2)
}
