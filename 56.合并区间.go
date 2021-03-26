/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
package main

/*
// 太慢了
import "sort"

type sortedIntervals [][]int
func (s sortedIntervals) Len() int { return len(s) }
func (s sortedIntervals) Less(i,j int) bool {
	return s[i][1] < s[j][1] || (s[i][1] == s[j][1] && s[i][0] < s[j][0])
}
func (s sortedIntervals) Swap(i,j int) { s[i],s[j] = s[j],s[i] }

func merge(intervals [][]int) [][]int {
	sortedintervals := sortedIntervals(intervals)
	sort.Sort(sortedintervals)
	l := len(intervals)
	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}
	for i := 0 ; i < l ; i++ {
		for j := 0 ; j < i ; j++ {
			if sortedintervals[i][0] <= sortedintervals[j][1] {
				sortedintervals[i][0] = min(sortedintervals[j][0],sortedintervals[i][0])
				copy(sortedintervals[j:],sortedintervals[i:])
				l -= i-j
				i = j
				// i++,则出去后i=j+1
			}
		}
	}
	return sortedintervals[:l]
}
*/
import "sort"

type sortedIntervals [][]int

func (s sortedIntervals) Len() int { return len(s) }
func (s sortedIntervals) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}
func (s sortedIntervals) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func merge(intervals [][]int) [][]int {
	sortedintervals := sortedIntervals(intervals)
	sort.Sort(sortedintervals)
	ans := [][]int{}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for _,interval := range intervals {
		l,r := interval[0],interval[1]
		if len(ans) == 0 || ans[len(ans)-1][1] < l {
			ans = append(ans,interval)
		}else{
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1],r)
		}
	}
	return ans
}

// @lc code=end
