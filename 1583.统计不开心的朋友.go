/*
 * @lc app=leetcode.cn id=1583 lang=golang
 *
 * [1583] 统计不开心的朋友
 */

// @lc code=start
package main

func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {
	// _pairs[i]记录的是i的配对
	_pairs := make([]int,n)
	// _preferences[i][j]记录的是j在i的朋友列表中的亲近程度
	_preferences := make([][]int,n)
	for i := range preferences {
		_preferences[i] = make([]int,n)
		for j := range preferences[i] {
			_preferences[i][preferences[i][j]] = j
		}
	}
	for i := range pairs {
		first,second := pairs[i][0],pairs[i][1]
		_pairs[first] = second
		_pairs[second] = first
	}

	result := 0
	// 为了便于理解，分别用x,y,u,v来表示，接近题意
	for x := range _pairs {
		y := _pairs[x]
		idxY := _preferences[x][y]
		for j := 0 ; j < idxY ; j++ {
			u := preferences[x][j]
			// xu的亲近程度高于xy
			v := _pairs[u]
			if _preferences[u][x] < _preferences[u][v] {
				// ux的亲近程度高于uv
				result++
				break
			}
		}
	}
	return result
}
// @lc code=end

