/*
 * @lc app=leetcode.cn id=275 lang=golang
 *
 * [275] H 指数 II
 */

// @lc code=start
func hIndex(citations []int) int {
	n := len(citations)
	l,r := 0,n-1
	for l < r {
		mid := (l+r)/2
		if citations[mid] >= n-mid {
			r = mid
		}else{
			l = mid+1
		}
	}
	if citations[r] == 0 {
		return 0
	}
	return n-r
}
// @lc code=end

