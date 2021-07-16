package main
import "sort"
func search(nums []int, target int) int {
	/*
	if len(nums) == 0 {
		return 0
	}
	var low int
	if nums[0] == target {
		low = 0
	}else{
		low = sort.Search(len(nums),func(i int) bool {
			return nums[i] == target && nums[i-1] != target
		})
	}
	if low == len(nums) {
		return 0
	}
	high := sort.Search(len(nums),func(i int) bool {return nums[i] > target})
	return high-low
	*/
	high := sort.Search(len(nums),func(i int) bool {return nums[i] > target})
	var low int
	if high == 0 {
		return 0
	}else{
		low = high
		for i := high-1 ; i >= 0 ; i-- {
			if nums[i] == target {
				low = i
			}
		}
	}
	return high-low
}