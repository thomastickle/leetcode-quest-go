package main

import "fmt"

func getConcatenation(nums []int) []int {
	return append(nums, nums...)
}

func shuffle(nums []int, n int) []int {
	var output = make([]int, len(nums))
	for i := range n {
		output[i*2], output[i*2+1] = nums[i], nums[i+n]
	}
	return output
}

func main() {
	fmt.Println("Hello Go World!")
}
