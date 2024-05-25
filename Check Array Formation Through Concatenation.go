package leetcodetest

func canFormArray(arr []int, pieces [][]int) bool {
	for _, piece := range pieces {
		aindex := -1
		for i, v := range arr {
			if v == piece[0] {
				aindex = i
			}
		}
		if aindex == -1 {
			return false
		}
		if len(piece) < 2 {
			continue
		}
		for _, v := range piece {
			if aindex < len(arr) && arr[aindex] == v {
				aindex++
			} else {
				return false
			}
		}
	}
	return true
}
