package leetcodetest

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	var n uint64
	if num < 0 {
		n = 1<<32 + uint64(num)
	} else {
		n = uint64(num)
	}
	hex := "0123456789abcdef"
	var s string
	for n > 0 {
		a := n % 16
		s = string(hex[a]) + s
		n /= 16
	}
	return s
}
