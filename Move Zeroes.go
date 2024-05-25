package leetcodetest

func moveZeroes(nums []int) {
	zCount, idx := 0, 0
	for _, v := range nums {
		if v == 0 {
			zCount++
		} else {
			nums[idx] = v
			idx++
		}
	}
	for i := len(nums) - 1; i >= 0 && zCount > 0; i-- {
		nums[i] = 0
		zCount--
	}
}
