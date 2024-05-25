package leetcodetest

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
		if maxArea < area {
			maxArea = area
		}
	}
	return maxArea
}
