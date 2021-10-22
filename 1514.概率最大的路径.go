/*
 * @lc app=leetcode.cn id=1514 lang=golang
 *
 * [1514] 概率最大的路径
 */

// @lc code=start
package main
func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	// w[i]记录以i为起点的边，map中key为终点，value为概率值
	w := make([]map[int]float64,n)
	for i,edge := range edges {
		if w[edge[0]] == nil {
			w[edge[0]] = map[int]float64{}
		}
		w[edge[0]][edge[1]] = succProb[i]
		if w[edge[1]] == nil {
			w[edge[1]] = map[int]float64{}
		}
		w[edge[1]][edge[0]] = succProb[i]
	}

	distance := make([]float64,n)
	for k,v := range w[start] {
		distance[k] = v
	}
	// if distance[end] != 0 {
	// 	return distance[end]
	// }

	notseen := map[int]bool{}
	for i := 0 ; i < n ; i++ {
		if i != start {
			notseen[i] = true
		}
	}
	for len(notseen) > 0 {
		// maxV来记录当前概率值最大的点
		maxV,maxDistance := -1,0.0
		for v := range notseen {
			if distance[v] > maxDistance {
				maxDistance = distance[v]
				maxV = v
			}
		}

		if maxV == end {
			return maxDistance
		}
		if maxV == -1 {
			break
		}
		for next,prob := range w[maxV] {
			if distance[maxV] * prob > distance[next] {
				distance[next] = distance[maxV] * prob
			}
		}
		delete(notseen,maxV)
	}
	return distance[end]
}
// @lc code=end

