package leetcodetest

import "strings"

func bucketSort(s string) string {
	list := make(map[rune]int)
	count := make([][]rune, len(s)+1)
	for _, v := range s {
		list[v]++
	}
	for v, time := range list {
		count[time] = append(count[time], v)
	}
	var result string
	for i := len(s); i >= 0; i-- {
		carr := count[i]
		for _, v := range carr {
			result += strings.Repeat(string(v), i)
		}
	}
	return result
}

func frequencySort(s string) string {
	return bucketSort(s)
}
