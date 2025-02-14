/*
 * @lc app=leetcode.cn id=1552 lang=golang
 *
 * [1552] 两球之间的磁力
 */

// @lc code=start
// 最大化最小值，二分查找
func maxDistance(position []int, m int) int {
	sort.Ints(position)
    l, r := 0, (position[len(position)-1] - position[0])
	res := 0
	for (l <= r) {
		mid := l + (r-l)/2
		if check(position, mid, m) {
			res = mid
			l = mid + 1 
		}else{
			r = mid - 1
		}
	}
	return res
}

func check(position []int, mid int, m int) bool {
	pre := position[0]
	m--
	for i := 1; i < len(position); i++ {
		if position[i] - pre >= mid {
			pre = position[i]
			m--
			if m == 0 {
				return true
			}
		}
	}
	return false
}
// @lc code=end

