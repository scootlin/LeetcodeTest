package leetcodetest

import "fmt"

func BinSearch(nums []int, target int) int {
	start := 0
	numsl := len(nums)
	end := numsl - 1
	i := numsl / 2
	for start <= end {
		fmt.Println(start, i, end)
		if nums[i] == target {
			return i
		} else if nums[i] > target {
			end = i - 1
		} else if nums[i] < target {
			start = i + 1
		}
		i = (end + start) / 2
	}
	return -1
}
