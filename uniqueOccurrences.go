package leetcodetest

import (
	"sort"
)

func uniqueOccurrences(arr []int) bool {
	sort.Ints(arr)
	count := make(map[int]int)
	val := make(map[int]int)
	prev := arr[0]
	c := 0
	for _, v := range arr {
		val[v]++
		if v != prev {
			c = val[prev]
			if _, ok := count[c]; ok {
				return false
			} else {
				count[c] = prev
			}
		}
		prev = v
	}
	if _, ok := count[val[arr[len(arr)-1]]]; ok {
		return false
	}

	return true
}
