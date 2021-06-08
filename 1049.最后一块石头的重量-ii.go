/*
 * @lc app=leetcode.cn id=1049 lang=golang
 *
 * [1049] 最后一块石头的重量 II
 */

// @lc code=start
package main
import "math"



func lastStoneWeightII(stones []int) int {
	var sum int
	for _,stone := range stones {
		sum += stone
	}
	t := sum/2
	// 思路是这样的，我们把所有数分成两堆，两堆和的差值要最小
	// 那么每一堆的总和应该越接近sum的一半
	// 这也是我们把总和上界定为sum/2的原因

	// 若常规建一个二维数组，f[i][j]表示组合前i项，得到的总和不超过j的最大值
	// f[j]这个是滚动数组用法
	f := make([]int,t+1)
	n := len(stones)
	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := 1 ; i <= n ; i++ {
		x := stones[i-1]
		f1 := make([]int,t+1)
		for j := 0 ; j <= t ; j++ {
			f1[j] = f[j]
			if j >= x {
				f1[j] = max(f1[j],f[j-x]+x)
			}
		}
		copy(f,f1)
	}
	return int(math.Abs(float64(sum-2*f[t])))
}
// @lc code=end

