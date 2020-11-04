package LeetcodeTest

import (
	"fmt"
	"strconv"
)

func BinaryGape(N int) int {
	str := ""
	max := 0
	left := 0
	right := 0
	index := 0
	arr := make(map[int]int)
	if N&(N-1) == 0 {
		return 0
	}
	for N > 0 {
		tmp := N % 2
		arr[index] = tmp
		// fmt.Println(tmp)
		N = N / 2
		str = strconv.Itoa(tmp) + str
		if tmp == 1 {
			if left <= right {
				left = right
				right = index
			} else {
				left = index
			}
			// fmt.Printf("left = %d, right = %d\n", left, right)
			gape := right - left
			if gape > 2 && gape > max && (arr[left] == 1 && arr[right] == 1) {
				fmt.Println(gape)
				max = gape
			}
		}

		index++
	}
	if max == 0 {
		return max
	} else {
		return max - 1
	}
}
