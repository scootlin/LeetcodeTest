package leetcodetest

func lengthOfLastWord(s string) int {
	start := 0
	end := len(s) - 1
	slen := 0
	for start <= end {
		if s[end] == 32 {
			end--
			continue
		}
		str := s[start]
		if str != 32 {
			slen++
		} else {
			slen = 0
		}
		start++
	}
	return slen
}
