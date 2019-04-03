package main

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	if totalLen%2 != 0 {
		return float64(nums[totalLen/2])
	} else {
		n1, n2 := nums[totalLen/2], nums[(totalLen/2)-1]
		if n1 == n2 {
			return float64(n1)
		} else {
			return float64(n1+n2) / 2
		}
	}
}
