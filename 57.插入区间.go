/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 */

// @lc code=start
package main
func insert(intervals [][]int, newInterval []int) [][]int {
	left,right := newInterval[0],newInterval[1]
	merged := false
	ans := [][]int{}
	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}
	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}
	for _,interval := range intervals {
		if interval[0] > right {
			if !merged {
				// ans = append(ans,newInterval)
				// 不写上面这句是因为可能已经经过合并了，left和right更新
				ans = append(ans,[]int{left,right})
				merged = true
			}
			ans = append(ans,interval)
		}else if interval[1] < left {
			ans = append(ans,interval)
		}else{
			// 有交集就计算并集
			left = min(left,interval[0])
			right = max(right,interval[1])
		}
	}
	if !merged {
		ans = append(ans,[]int{left,right})
	}
	return ans
}

// @lc code=end
