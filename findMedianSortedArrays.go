package leetcodetest

import "log"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i := len(nums1)
	j := len(nums2)
	total := i + j
	median := (total) / 2
	if i <= 0 {
		return calc(nums2, median, total)
	}

	if j <= 0 {
		return calc(nums1, median, total)
	}

	m := (total) - 1
	mergeArray := make([]int, total)

	i--
	j--

	pivot := median
	if pivot%2 == 0 {
		pivot--
	}

	for m >= pivot {
		if i >= 0 && j >= 0 {
			if nums1[i] <= nums2[j] {
				mergeArray[m] = nums2[j]
				j--
			} else {
				mergeArray[m] = nums1[i]
				i--
			}
		} else if i >= 0 {
			mergeArray[m] = nums1[i]
		} else if j >= 0 {
			mergeArray[m] = nums2[i]
		}
		m--
	}

	// log.Println(mergeArray, median, pivot)
	return calc(mergeArray, median, total)
}

func calc(arr []int, m, t int) float64 {
	if m%2 != 0 || (m) <= 2 {
		return float64(arr[m])
	} else {
		log.Println(float64((arr[m] + arr[m-1])) / 2.0)
		return float64((arr[m] + arr[m-1])) / 2.0
	}
}
