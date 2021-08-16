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

func removeDuplicates(nums []int) int {
	var res []int
	for i := 0; i <= len(nums)-1; i++ {
		if contains(res, nums[i]) {
			continue
		} else {
			res = append(res, nums[i])
		}
	}
	return len(res)
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	num := removeDuplicates(nums)
	fmt.Println(num)
}
