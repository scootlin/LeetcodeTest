package leetcodetest

func maxSubArray(nums []int) int {
	if len(nums) <= 1 {
		return nums[0]
	}
	max := -1 << 31
	num := 0
	lnums := len(nums)
	for i := 0; i < lnums; i++ {
		num += nums[i]
		if num < nums[i] {
			num = nums[i]
		}
		if max < num {
			max = num
		}
	}
	return max
}
