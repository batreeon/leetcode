/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
package main
func trap(height []int) int {

	/*
	// 直接用图形
	// 所有的h1,h2相加，结果等于，所有的雨水加所有的桶加最高*宽度（左右拼起来）
	// 然后每一步减去height[i]减去所有的桶
	// 最后返回结果时再减去多余的最高*宽
	result,h1,h2 := 0,0,0
	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}
	n := len(height)
	for i := 0 ; i < n ; i++ {
		h1 = max(h1,height[i])
		h2 = max(h2,height[n-1-i])
		result += h1 + h2 - height[i]
	}
	return result - n*h1
	*/

	// 单调栈
	// 存下标
	stack := []int{}
	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}
	result := 0
	for i,h := range height {
		for n := len(stack) ; n > 0 && h > height[stack[n-1]] ; {
			top := stack[n-1]
			stack = stack[:n-1]
			n = len(stack)
			if n == 0 {
				// 左边没有更高的，存不住水，出现在最右边把
				break
			}
			left := stack[n-1]
			wide := i-left-1
			hheight := min(height[left],h) - height[top]
			result += wide*hheight
		}
		stack = append(stack,i)
	}
	return result
}
// @lc code=end

