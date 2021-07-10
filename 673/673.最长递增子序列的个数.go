/*
 * @lc app=leetcode.cn id=673 lang=golang
 *
 * [673] 最长递增子序列的个数
 */

// @lc code=start
package main
func findNumberOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	length := make([]int,len(nums))
	count := make([]int,len(nums))
	for i := range count {
		count[i] = 1
	}
	for i,num1 := range nums {
		for j,num2 := range nums[:i] {
			if num2 < num1 {
				if length[j] >= length[i] {
					length[i] = 1+length[j]
					count[i] = count[j]
				} else if length[j] + 1 == length[i] {
					// 这一步是怎么来的呢，比如说有序列1，3，5，4，7
					// 当nums[i] == 7时，当nums[j] = 5时，length[i]被更新为了length[j]+1,count[i]被更新为count[j]
					// 上面就相当于得到了1,3,5,7序列
					// 但其实也还可以得到1,3,4,7序列
					// 因此接下来当nums[j]==4时，通过4也可以构造出一个最长的序列（最长是局部最长，上一步得到的）
					// 所以就把1,3,4,7这一条路给加进来
					count[i] += count[j]
				}
			}
		}
	}
	maxlen := 0
	for _,leng := range length {
		if leng > maxlen {
			maxlen = leng
		}
	}
	result := 0
	for i,leng := range length {
		if leng == maxlen {
			result += count[i]
		}
	}
	return result
}
// @lc code=end

