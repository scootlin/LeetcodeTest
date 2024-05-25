package leetcodetest

func removeDuplicates(nums []int) int {
	// list := make(map[int]int)
	if len(nums) <= 1 {
		return len(nums)
	}
	// nlen := 0
	i := 0
	for i < len(nums)-1 {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)
}
