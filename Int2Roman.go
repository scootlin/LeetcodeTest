package leetcodetest

import (
	"math"
)

var intToRomanMap = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

func intToRoman(num int) string {
	if str, ok := intToRomanMap[num]; ok {
		return str
	}
	power := 0
	remainder := 0
	output := ""
	for num >= 10 {
		remainder = num % 10
		num = num / 10
		// fmt.Println("quotient = ", num, "power = ", power, "remainder = ", remainder)
		if remainder != 0 {
			partstr := getRoman(power, remainder)
			output = partstr + output
			// fmt.Println("output = ", partstr)
		}
		power++
	}
	last := num * int(math.Pow(10, float64(power)))
	if last >= 1000 {
		// fmt.Println("M ====== ", num)
		for i := num; i > 0; i-- {
			output = "M" + output
		}
	} else {
		partstr := getRoman(power, num)
		// fmt.Println("quotient = ", num, "power = ", power, "remainder = ", remainder, "output = ", partstr)
		output = partstr + output
	}
	return output
}

func getRoman(power int, remainder int) string {
	median := 5 * int(math.Pow(10, float64(power))) //5,50,500
	span := 1 * int(math.Pow(10, float64(power)))   //1,10,100
	partstr := ""
	if remainder <= 5 {
		// fmt.Println(" <<< == 5")
		partstr = conver2Raman(median, span, 5, remainder)
	} else if remainder > 5 && remainder < 9 {
		// fmt.Println("6~~~~~9")
		partstr = conver2Raman(median, span, remainder, 5)
	} else {
		// fmt.Println(">>>>> == 9")
		partstr = conver2Raman(median*2, span, 10, remainder)
	}
	return partstr
}

func conver2Raman(median, span int, max int, target int) string {
	base := intToRomanMap[median]
	// fmt.Println("Base = ", base)
	minus := intToRomanMap[span]
	// fmt.Println("minus = ", minus)
	if target < 4 {
		base = ""
		for i := target; i > 0; i-- {
			base += minus
		}
	} else if target < max && target != 9 && target != 4 {
		len := max - target
		for i := len; i > 0; i-- {
			base += minus
		}
	} else if target%5 == 0 {
		return base
	} else {
		minus := intToRomanMap[span]
		base = minus + base
	}
	return base
}
