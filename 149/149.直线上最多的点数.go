/*
 * @lc app=leetcode.cn id=149 lang=golang
 *
 * [149] 直线上最多的点数
 */

// @lc code=start
package main
func maxPoints(points [][]int) int {
	// 我人傻了，直接暴力啊，以每个点为起点计算与其他点的斜率
	// 这样i第一次取到一条直线上的某个点时，就能得到这一条线上的所有点数
	// 因此每个i不需要在将所有点遍历一遍。这就是j为什么从i+1开始取
	if len(points) == 1 {
		return 1
	}
	result := 0
	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := 0 ; i < len(points) ; i++ {
		k := map[float64]int{}//存放斜率
		_k := map[float64]int{}//用来记录斜率不存在时的情况
		pi := points[i]
		for j := i+1 ; j < len(points) ; j++ {
			pj := points[j]
			if pi[0] != pj[0] {
				kij := (float64(pi[1])-float64(pj[1]))/(float64(pi[0])-float64(pj[0]))
				if _,ok := k[kij] ; !ok {
					k[kij] = 2
				}else{
					k[kij]++
				}
				result = max(result,k[kij])
			}else{
				if _,ok := _k[float64(pi[0])] ; !ok {
					_k[float64(pi[0])] = 2
				}else{
					_k[float64(pi[0])]++
				}
				result = max(result,_k[float64(pi[0])])
			}
		}
	}
	return result
}
// @lc code=end

