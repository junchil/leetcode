package main

import "fmt"

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func intersection(nums1 []int, nums2 []int) []int {
	var res []int
	for _, num := range nums1 {
		if contains(nums2, num) && !contains(res, num) {
			res = append(res, num)
		}
	}
	return res
}

func main() {
	nums1 := []int{4, 9, 2, 2, 1}
	nums2 := []int{2, 2, 4, 9}
	res := intersection(nums1, nums2)
	fmt.Println(res)
}
