/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
package main
func searchRange(nums []int, target int) []int {
	// 找到任一个target元素的下标，若不存在就返回-1
	findTarget := func(nums *[]int,target int) int {
		l,r := 0,len(*nums)-1
		for r >= l {
			m := (l+r)/2
			if (*nums)[m] == target {
				return m
			}else if (*nums)[m] > target {
				r = m-1
			}else if (*nums)[m] < target {
				l = m+1
			}
		}
		return -1
	}
	// 在上一个函数的返回结果index的左边找左边界
	findLeft := func(nums *[]int,target,l,r int) int {
		for {
			mid := (l+r)/2
			if (*nums)[mid] == target {
				if mid == 0 {
					return mid
				}else if (*nums)[mid-1] != target {
					return mid
				}else{
					r = mid-1
				}
			}else{
				if (*nums)[mid+1] == target {
					return mid+1
				}else{
					l = mid+1
				}
			}
		}
	}
	// 在上一个函数的返回结果index的右边寻找右边界
	findRight := func(nums *[]int,target,l,r int) int {
		for {
			mid := (l+r)/2
			if (*nums)[mid] == target {
				if mid == len(*nums)-1 {
					return mid
				}else if (*nums)[mid+1] != target {
					return mid
				}else{
					l = mid+1
				}
			}else{
				if (*nums)[mid-1] == target {
					return mid-1
				}else{
					r = mid+1
				}
			}
		}
	}
	
	left,right := -1,-1
	length := len(nums)
	index := findTarget(&nums,target)
	if index != -1 {
		if index == 0 {
			left = index
		}else if nums[index-1] != target {
			left = index
		}else{
			left = findLeft(&nums,target,0,index-1)
		}
		
		if index == length-1 {
			right = index
		}else if nums[index+1] != target {
			right = index
		}else{
			right = findRight(&nums,target,index+1,length-1)
		}
	}
	return []int{left,right}
}
// @lc code=end

