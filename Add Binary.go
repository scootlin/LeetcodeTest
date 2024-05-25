package leetcodetest

func addBinary(a string, b string) string {
	shift := false
	//strA := ""
	//strB := ""
	sum := ""
	aLen := len(a)
	bLen := len(b)
	start := 1
	var chA, chB string

	for aLen-start >= 0 || bLen-start >= 0 {
		// fmt.Println(start)

		if aLen-start >= 0 {
			chA = string(a[aLen-start])
		} else {
			chA = "0"
		}
		if bLen-start >= 0 {
			chB = string(b[bLen-start])
		} else {
			chB = "0"
		}
		//  fmt.Println("chA,chB,shift",chA,chB,shift)
		if chA == "1" && chB == "1" {
			if shift == true {
				sum = "1" + sum
			} else {
				sum = "0" + sum
			}
			shift = true
		} else if chA == "0" && chB == "0" {
			if shift == true {
				sum = "1" + sum
			} else {
				sum = "0" + sum
			}
			shift = false
		} else {
			if shift == true {
				sum = "0" + sum
				shift = true
			} else {
				sum = "1" + sum
				shift = false
			}
		}
		start++
	}
	if shift {
		sum = "1" + sum
	}
	return sum
}
