package main

import "fmt"

func getConcatenation(nums []int) []int {
	return append(nums, nums...)
}

func shuffle(nums []int, n int) []int {
	var output = make([]int, len(nums))
	for i := 0; i < n; i++ {
		output[i*2], output[i*2+1] = nums[i], nums[i+n]
	}

	return output
}

func main() {
	var input = []int{2, 5, 1, 3, 4, 7}
	var output = shuffle(input, 3)
	fmt.Println(output)
	fmt.Println("Hello Go World!")
}
