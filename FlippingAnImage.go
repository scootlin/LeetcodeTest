package main

import (
	"fmt"
)

func main() {
	fmt.Println("Flip matrix  : ", flipAndInvertImage([][]int{[]int{0, 0, 1}, []int{1, 1, 0}, []int{1, 0, 0}}))
}
func flipAndInvertImage(A [][]int) [][]int {
	for _, row := range A {
		last := len(row) - 1
		for j, _ := range row {
			if j < last-1 {
				row[j], row[last-j] = row[last-j], row[j]
			}
			if row[j] == 0 {
				row[j] = 1
			} else {
				row[j] = 0
			}
		}
	}
	return A
}
