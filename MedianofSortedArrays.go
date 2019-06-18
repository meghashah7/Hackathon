package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 3}

	nums2 := []int{2}

	fmt.Println(findMedianSortedArrays(nums1, nums2))

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	var median float64

	nums1 = append(nums1, nums2...)

	sort.Ints(nums1)
	fmt.Println(nums1)
	length := len(nums1)
	fmt.Println(length)
	half := length / 2
	if length%2 == 0 {
		median = float64(float64(nums1[half-1]+nums1[half]) / 2)
	} else {
		median = float64(nums1[half])
	}
	return median

}
