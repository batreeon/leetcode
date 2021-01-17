/*
 * @lc app=leetcode.cn id=1365 lang=golang
 *
 * [1365] 有多少小于当前数字的数字
 */

// @lc code=start
// package main

// import (
// 	"fmt"
// 	"sort"
// )

func smallerNumbersThanCurrent(nums []int) []int {
	nummap := make(map[int]int)
	//sortnums := nums//错误用法,这样传递的是指针
	sortnums := make([]int, len(nums))
	copy(sortnums, nums)
	sort.Ints(sortnums)
	// fmt.Println(nums)
	// fmt.Println(sortnums)
	i := 0
	for ; sortnums[i] != sortnums[len(sortnums)-1]; i++ {
		nummap[sortnums[i]] = i
		// fmt.Println(sortnums[i], i)
		for ; sortnums[i+1] == sortnums[i]; i++ {
		}
		// fmt.Println(i)
	}
	nummap[sortnums[i]] = i
	// fmt.Println(sortnums[i], i)
	for i, v := range nums {
		nums[i] = nummap[v]
		// fmt.Println(i, v, nummap[v])
	}
	return nums
}

// func main() {
// 	ans := smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3})
// 	fmt.Println(ans)
// }

// @lc code=end
