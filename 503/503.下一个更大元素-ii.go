/*
 * @lc app=leetcode.cn id=503 lang=golang
 *
 * [503] 下一个更大元素 II
 */

// @lc code=start
package main
func nextGreaterElements(nums []int) []int {
	/*
	//暴力法
	l := len(nums)
	ans := make([]int,l)
	for i := 0 ; i < l ; i++ {
		now := nums[i]
		j := (i+1)%l
		for ; j != i ; j = (j+1)%l {
			if nums[j] > now {
				break
			}
		}
		if j == i {
			ans[i] = -1
		}else{
			ans[i] = nums[j]
		}
	}
	return ans
	*/

	//单调栈
	l := len(nums)
	ans := make([]int,l)
	copy(ans,nums)
	//可以创建一个二维数组，记录值和下标，但我们选择只记下标，从nums读值
	stack := []int{}
	sl := 0//减少使用len(stack)的消耗
	for i := 0 ; i < 2 * l; i++ {
		j := i % l
		if ans[j] == nums[j] {
			ans[j] = -1
		}
		for sl > 0 && nums[stack[sl-1]] < nums[j] {
			ans[stack[sl-1]] = nums[j]
			stack = stack[:sl-1]//只有将要出栈的元素才是找到了下一个更大值的
			sl--
		}
		stack = append(stack,j)
		sl++
	}
	return ans
}
// @lc code=end

