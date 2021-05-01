/*
 * @lc app=leetcode.cn id=403 lang=golang
 *
 * [403] 青蛙过河
 */

// @lc code=start
package main
func canCross(stones []int) bool {
	// 深度优先遍历，k是上一步跳的步数，i是当前所在的石子下标
	visited := map[int]bool{}
	var dfs func(k,i int) bool
	dfs = func(k,i int) bool {
		// 遇到相同的子问题就提前退出，哇哇哇哇，这个map的构造好叼，学到了
		if visited[i*2000+k] {
			return false
		}else{
			visited[i*2000+k] = true
		}
		if i == len(stones)-1 {
			return true
		}
		for j := 1 ; i+j < len(stones) && stones[i+j] - stones[i] <= k+1 ; j++ {
			if stones[i+j] - stones[i] == k-1 {
				if dfs(k-1,i+j) {
					return true
				}
			}
			if stones[i+j] - stones[i] == k {
				if dfs(k,i+j) {
					return true
				}
			}
			if stones[i+j] - stones[i] == k+1 {
				if dfs(k+1,i+j) {
					return true
				}
			}
		}
		return false
	}
	if stones[1] - stones[0] != 1 {
		return false
	}
	return dfs(1,1)
}
// @lc code=end

