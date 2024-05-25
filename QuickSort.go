package leetcodetest

import (
	"math/rand"
	"time"
)

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(9999999) - rand.Intn(9999999)
	}
	return slice
}

func quickSort(ints []int) []int {
	if len(ints) < 2 {
		return ints
	}
	// right := len(ints)-1
	middle := (len(ints) - 1) / 2
	// poivot := ints[len(ints)] - 1
	// ints[right],ints[middle] = ints[middle],ints[right]
	left := []int{}
	right := []int{}
	mount := []int{}
	for i, v := range ints {
		if i == middle {
			continue
		}
		if v > ints[middle] {
			right = append(right, v)
		} else {
			left = append(left, v)
		}
	}
	mount = append(mount, quickSort(left)...)
	mount = append(mount, ints[middle])
	mount = append(mount, quickSort(right)...)
	return mount
}

func Inplace(ints []int) []int {
	if len(ints) < 2 {
		return ints
	}
	pivot := rand.Int() % len(ints)
	left, right := 0, len(ints)-1
	ints[pivot], ints[right] = ints[right], ints[pivot] //Replace pivot index to right
	for i, val := range ints {
		if val < ints[right] {
			ints[i], ints[left] = ints[left], ints[i]
			left++
		}
	}
	ints[left], ints[right] = ints[right], ints[left]
	Inplace(ints[:left])
	Inplace(ints[left+1:])
	return ints
}
