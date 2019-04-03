package main

import "fmt"

func main() {
	//fmt.Println(isPalindrome(100110011001))
	fmt.Println("Max Area = ", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
func maxArea(height []int) int {
	maxArea := 0
	left := 0
	right := len(height) - 1
	for right > left {
		var x, y int
		var area int
		x = right - left
		if height[left] < height[right] {
			y = height[left]
			left++
		} else {
			y = height[right]
			right--
		}
		area = x * y
		fmt.Println("Area = ", area)
		if maxArea < area {
			maxArea = area
		}
	}
	return maxArea
}
