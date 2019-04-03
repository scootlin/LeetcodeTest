package main

import (
	"fmt"
)

func main() {
	// groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Println("Two sum  : ", twoSum([]int{3, 2, 4}, 6))
}
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
