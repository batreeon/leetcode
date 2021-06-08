/*
 * @lc app=leetcode.cn id=403 lang=golang
 *
 * [403] 青蛙过河
 */

// @lc code=start
func canCross(stones []int) bool {

	/*
	// dfs+记忆化
	// 之前做过，但用的dfs,这次用其他方法试试
	// 记录石子的位置
	m := make(map[int]bool)
	// 记忆化
	cache := make(map[int]bool)

	// 参数分别是当前所在位置和上一步跳跃步数
	var dfs func(i,k int) bool
	dfs = func(i,k int) bool {
		if i == stones[len(stones)-1] {
			return true
		}
		if v,ok := cache[i*2000+k] ; ok {
			return v
		}
		for j := -1 ; j <= 1 ; j++ {
			if k+j <= 0 {
				// 不能后退
				continue
			}
			nextIdx := i+k+j
			if m[nextIdx] {
				ans := dfs(nextIdx,k+j)
				if ans {
					return true
				}
			}
		}
		cache[i*2000+k] = false
		return false
	}

	for _,stone := range stones {
		m[stone] = true
	}
	if !m[1] {
		return false
	}
	return dfs(1,1)
	*/

	// 动规解法
	if stones[1] != 1 {
		return false
	}
	n := len(stones)
	// f[i][j]代表当跳跃j步到达下标i时，最终是否能到达终点
	f := make([][]bool,n+1)
	f[0] = make([]bool,n+1)
	f[1] = make([]bool,n+1)
	f[1][1] = true
	for i := 2 ; i < n ; i++ {
		f[i] = make([]bool,n+1)
		// 上一步所在的下标j，从j跳到i
		for j := 1 ; j < i ; j++ {
			// 步数k,由j跳到i需要的步数
			k := stones[i]-stones[j]
			// 这一步的结果由上一步的状态决定

			// 下面这一步没搞懂
			// if k+1 < n+1 {
			if k <= j+1 {
				f[i][k] = f[j][k-1] || f[j][k] || f[j][k+1]
			}
		}
	}
	for i := 1 ; i < n ; i++ {
		if f[n-1][i] {
			return true
		}
	}
	return false
}
// @lc code=end

