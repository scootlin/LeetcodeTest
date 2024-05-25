package leetcodetest

func twoSum(nums []int, target int) []int {
	tmpMap := make(map[int]int)
	for i, num := range nums {
		if _, ok := tmpMap[target-num]; ok {
			return []int{tmpMap[target-num], i}
		}
		tmpMap[num] = i
	}
	return []int{}
}
