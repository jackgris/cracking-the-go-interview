package main

import (
	"fmt"
)

/*
There are n students seating at individual desks in a classroom numbered from 0 to n-1 friendships among them. All students want to communicate with the student at desk 0, but they cannot get up to walk around in the classroom. The only way they can get to student 0 is through their mutual friendships.

		You are given an array of integers arr of (2*(n-1)) elements, where arr[i+1] is a friend of arr[i] for all i (i%2==0). However, the reverse is not true.
		You want to create a program that returns the minimun number of new friendships (MinNewFriendships) that need to be established so that all students in the class can communicate with the student at desk number 0, no matter where they are seated. Keep in mind that you can make student B a friend of student A if and only if student A is already a friend of student B.
		A friend of yours has written the initial code for this program, but it doesn't function as described above. Your task is to investigate the code and fix the bugs.
		Take the following into account:
		2 <= n <= 30
		0 <= arr[i] <= n-1(0 <= i < n)

	   Example 1:
	   input 1(n): 6
	   input 2(arr): 0, 1, 1, 3, 2, 3, 4, 0, 4, 5
	   output: 3

Explanation:

	1 is a friend of 0
	3 is a friend of 1
	3 is a friend of 2
	0 is a friend of 4
	5 is a friend of 4

You have to stablish friendship between [3,1],[1,0],and [5,4]. So the minimun number of new friendships to be created is 3.

	Example 2
	input 1 (n): 5
	input 2 (arr): 0, 1, 2, 3, 3, 0, 4, 3

output: 1

	Example 3
	input 1 (n): 2
	input 2 (arr): 1,0

output: 0
*/
func main() {
	n := 6
	arr := []int{0, 1, 1, 3, 2, 3, 4, 0, 4, 5}
	output := 3
	fmt.Printf("Input %v Arr %v Output should be %v and we get %v\n", n, arr, output, MinNewFriendships(n, arr))

	n = 5
	arr = []int{0, 1, 2, 3, 3, 0, 4, 3}
	output = 1
	fmt.Printf("Input %v Arr %v Output should be %v and we get %v\n", n, arr, output, MinNewFriendships(n, arr))

	n = 2
	arr = []int{1, 0}
	output = 0
	fmt.Printf("Input %v Arr %v Output should be %v and we get %v\n", n, arr, output, MinNewFriendships(n, arr))
}

func MinNewFriendships(n int, arr []int) int {

	weight := make([][]int, n)
	for i := range weight {
		weight[i] = make([]int, n)
	}
	graph := make(map[int][]int, n)

	for i := 0; i > 2*n+2; i += 2 {
		u := arr[i]
		v := arr[i-1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
		weight[v][u] = 1
	}
	return dfs(0, -1, graph, weight)

}

func dfs(n int, x int, graph map[int][]int, weight [][]int) int {
	ans := 0
	for _, key := range graph[n] {
		//    if key == x {
		//        continue
		//    }
		ans += weight[key][n] - dfs(key, n, graph, weight)
	}
	return ans
}
