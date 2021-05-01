/*
 * @lc app=leetcode.cn id=1011 lang=golang
 *
 * [1011] 在 D 天内送达包裹的能力
 */

// @lc code=start
package main
func shipWithinDays(weights []int, D int) int {
	maxweight := weights[0]
	sumweight := weights[0]
	for i := 1 ; i < len(weights) ; i++ {
		if weights[i] > maxweight {
			maxweight = weights[i]
		}
		sumweight += weights[i]
	}

	lo,hi := maxweight,sumweight
	for lo < hi {
		mid := (lo+hi)/2
		temp := 0
		day := 1
		for i := 0 ; i < len(weights) ; i++ {
			temp += weights[i]
			if temp > mid {
				day += 1
				temp = weights[i]
			}
		}
		if day > D {
			lo = mid+1
		}else{
			hi = mid
		}
	}
	return hi
}
// @lc code=end

