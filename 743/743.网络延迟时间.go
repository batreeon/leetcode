/*
 * @lc app=leetcode.cn id=743 lang=golang
 *
 * [743] 网络延迟时间
 */

// @lc code=start
package main
import "math"
func networkDelayTime(times [][]int, n int, k int) int {
	es := make([][][]int,n+1)
	for _,time := range times {
		es[time[0]] = append(es[time[0]],[]int{time[1],time[2]})
	}

	distance := make([]int,n+1)
	for i := 1 ; i <= n ; i++ {
		distance[i] = -1
	}
	for i := 0 ; i < len(es[k]) ; i++ {
		distance[es[k][i][0]] = es[k][i][1]
	}
	distance[k] = 0

	notseen := map[int]bool{}
	for i := 1 ; i <= n ; i++ {
		if i != k {
			notseen[i] = true
		}
	}
	for {
		minV,minDistance := 0,math.MaxInt64
		for v := range notseen {
			if distance[v] != -1 {
				if distance[v] < minDistance {
					minDistance = distance[v]
					minV = v
				}
			}
		}

		// 有的点无法到达
		if minV == 0 && len(notseen) > 0 {
			return -1
		}
		for _,e := range es[minV] {
			if distance[e[0]] == -1 {
				distance[e[0]] = distance[minV] + e[1]
			}else{
				if distance[minV] + e[1] < distance[e[0]] {
					distance[e[0]] = distance[minV] + e[1]
				} 
			}
		}
		delete(notseen,minV)
		if len(notseen) == 0 {
			break
		}
	}
	result := 0
	for i := 1 ; i <= n ; i++ {
		if distance[i] > result {
			result = distance[i]
		}
	}
	return result
}
// @lc code=end

