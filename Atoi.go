package leetcodetest

import (
	"strconv"
	"strings"
)

func myAtoi(str string) int {
	const Max = 1<<31 - 1
	const Min = -Max - 1
	slen := len(str)
	mstr := ""
	var prev byte = 0
	for i := 0; i < slen; i++ {
		ch := str[i]
		if ch == 32 || ch == 43 || ch == 45 || (ch >= 48 && ch <= 57) {
			if (ch == 32 || ch == 43 || ch == 45) && i != 0 && (prev >= 48 && prev <= 57) {
				break
			}
			mstr += string(ch)
		} else {
			break
		}
		prev = ch
	}
	mstr = strings.TrimSpace(mstr)
	f, err := strconv.ParseFloat(mstr, 32/64)
	if err == nil {
		if f > Max {
			return Max
		} else if f < Min {
			return Min
		}
		return int(f)
	}
	return 0
}
