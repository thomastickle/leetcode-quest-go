package arrays

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

func findMaxConsecutiveOnes(nums []int) int {
	var maxCount = 0
	var currentCount = 0

	for _, x := range nums {
		if x == 1 {
			currentCount++
		} else {
			if currentCount > maxCount {
				maxCount = currentCount
			}
			currentCount = 0
		}
	}

	if currentCount > maxCount {
		maxCount = currentCount
	}

	return maxCount
}
