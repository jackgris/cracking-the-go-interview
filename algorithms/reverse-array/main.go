package main

import "fmt"

func main() {
	var scores = []int{60, 50, 95, 80, 70, 88, 95}
	fmt.Printf("List = %d\n", scores)
	fmt.Printf("List = %d\n", reverse(scores))
}

func reverse(list []int) []int {
	lenght := len(list) - 1
	middle := lenght / 2

	for i := 0; i < middle; i++ {
		tmp := list[i]
		list[i] = list[lenght-i]
		list[lenght-i] = tmp
	}

	return list
}
