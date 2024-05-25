package leetcodetest

import (
	"strings"
)

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	total := 0
	strarr := strings.Split(strings.ToUpper(s), "")
	strlen := len(strarr)
	i := strlen - 1
	for i >= 0 {
		val := roman[strarr[i]]
		if i-1 >= 0 {
			prev := roman[strarr[i-1]]
			if prev < val {
				val -= prev
				i--
			}
		}
		total += val
		i--
	}
	return total
}
