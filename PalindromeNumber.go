package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(100110011001))
}

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	bs := []byte(str)
	strlen := len(bs)
	for i := strlen; i > strlen/2; i-- {
		if bs[i-1] != bs[strlen-i] {
			return false
		}
	}
	return true
}
