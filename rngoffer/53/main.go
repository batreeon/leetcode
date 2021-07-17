package main

import (
	"fmt"
	"sort"
)

func findFirstIndex1(nums []int, target int) int {
	var low int
	if nums[0] == target {
		low = 0
	} else {
		low = sort.Search(len(nums), func(i int) bool {
			return nums[i] == target
		})
	}
	return low
}
func findFirstIndex2(nums []int, target int) int {
	low := sort.Search(len(nums), func(i int) bool { return nums[i] > target-1 })
	if low != len(nums) && nums[low] != target {
		low = len(nums)
	}
	return low
}
func main() {
	nums := []int{1, 2, 3, 3, 3, 3, 4, 5, 9}
	fmt.Println(findFirstIndex1(nums, 3))
	fmt.Println(findFirstIndex2(nums, 3))
	nums = []int{0, 1, 2, 3, 4, 4, 4}
	fmt.Println(findFirstIndex1(nums, 2))
	fmt.Println(findFirstIndex2(nums, 2))
}
