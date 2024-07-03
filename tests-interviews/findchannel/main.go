package main

import "fmt"

/*
   There are n channels on your TV numbered 1 to n. The channels are arranged in such a manner that if you move through the channels in order, after the nth channel you reach the first channel again.

   More formally, moving from the ith channel to the next channel brings you to the (i+1)th channel where 1 <= i < n, and moving from the nth channel to the next channel brings you to the 1st channel.

   Something has gone wrong with your TV and every time you move k number of channel (counting the channel from which you started at in k), the channel you land on disappears and you automatically see the next channel remains on your TV or untill you can no longer move k channels, in which case your TV stays stuck in the last channel in the process.

   Your task is to write a program that finds the channel (FindChannel) you can watch. If the only channel left on the TV is the last channel and it is the nth channel, return -1.

Note: You're always starting from channel 1.

   Take the following into account:
   1 <= n <= 100
   1 <= k <= n
*/

func main() {
	n := 5
	k := 2
	output := 3
	fmt.Printf("Number of channels %d, delete %d channels, output %d. Result is %d\n", n, k, output, FindChannel(n, k))

	n = 6
	k = 2
	output = 5
	fmt.Printf("Number of channels %d, delete %d channels, output %d. Result is %d\n", n, k, output, FindChannel(n, k))

	n = 7
	k = 3
	output = 4
	fmt.Printf("Number of channels %d, delete %d channels, output %d. Result is %d\n", n, k, output, FindChannel(n, k))
}

func FindChannel(n int, k int) int {
	// Handle k >= n case for early termination
	if k >= n {
		return -1
	}
	// Create a slice to represent the channels.
	channels := []int{}
	// Initialize channels from 1 to n
	for i := 1; i <= n; i++ {
		channels = append(channels, i)
	}

	pos := 0
	for len(channels) > 1 {
		pos = (pos + k - 1) % len(channels)
		// Remove the channel at the calculated position.
		channels = append(channels[:pos], channels[pos+1:]...)
	}
	// If the last remaining channel is n, return -1.
	if channels[0] == n {
		return -1
	}
	// Otherwise, return the last remaining channel.
	return channels[0]
}
