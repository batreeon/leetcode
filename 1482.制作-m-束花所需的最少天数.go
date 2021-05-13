/*
 * @lc app=leetcode.cn id=1482 lang=golang
 *
 * [1482] 制作 m 束花所需的最少天数
 */

// @lc code=start
func minDays(bloomDay []int, m int, k int) int {
	if m > len(bloomDay)/k {
		return -1
	}
	min,max := bloomDay[0],bloomDay[0]
	for i := 1 ; i < len(bloomDay) ; i++ {
		if bloomDay[i] < min {
			min = bloomDay[i]
		}
		if bloomDay[i] > max {
			max = bloomDay[i]
		}
	}
	for min < max {
		mid := (min+max)/2
		if can(bloomDay,m,k,mid) {
			max = mid
		}else{
			min = mid+1
		}
	}
	return min
}
func can(bloomDay []int,m int,k int,day int) bool {
	cnt := 0
	flowers := 0
	for i := 0 ; i < len(bloomDay) ; i++ {
		if bloomDay[i] <= day {
			flowers++
			if flowers == k {
				cnt++
				if cnt == m {
					return true
				}
				flowers = 0
			}
		}else{
			flowers = 0
		}
	}
	return false
}
// @lc code=end