/*
 * @lc app=leetcode.cn id=1743 lang=golang
 *
 * [1743] 从相邻元素对还原数组
 */

// @lc code=start
package main

func restoreArray(adjacentPairs [][]int) []int {
	/*
	if len(adjacentPairs) <= 1 {
		return adjacentPairs[0]
	}
	nodes := []int{}
	index := map[int]int{}
	for _,pair := range adjacentPairs {
		n1,n2 := pair[0],pair[1]
		if _,ok := index[n1] ; !ok {
			index[n1] = len(nodes)
			nodes = append(nodes,n1)
		}
		if _,ok := index[n2] ; !ok {
			index[n2] = len(nodes)
			nodes = append(nodes,n2)
		}
	}
	numNode := len(nodes)
	m := make([][]int,numNode)
	for i := range m {
		m[i] = make([]int,numNode)
		for j := range m[i] {
			m[i][j] = numNode
		}
	}
	for _,pair := range adjacentPairs {
		i1,i2 := index[pair[0]],index[pair[1]]
		m[i1][i2] = 1
		m[i2][i1] = 1
	}
	start,end := -1,-1
	for k := 0 ; k < numNode ; k++ {
		for i := 0 ; i < numNode ; i++ {
			for j := 0 ; j < numNode ; j++ {
				if m[i][k]+m[k][j] < m[i][j] {
					m[i][j] = m[i][k]+m[k][j]
					if m[i][j] == numNode-1 {
						start,end = i,j
					}
				}
			}
		}
	}
	result := make([]int,numNode)
	result[0],result[numNode-1] = nodes[start],nodes[end]
	
	for i := 1 ; i < numNode-1 ; i++ {
		for j := 0 ; j < numNode ; j++ {
			if m[start][j] == i && m[j][end] == numNode-1-i {
				result[i] = nodes[j]
				break
			}
		}
	}
	return result
	*/

	cnt := map[int]int{}
	neibour := map[int][]int{}
	for _,pair := range adjacentPairs {
		n1,n2 := pair[0],pair[1]
		cnt[n1]++
		cnt[n2]++
		if _,ok := neibour[n1] ; !ok {
			neibour[n1] = []int{}
		}
		neibour[n1] = append(neibour[n1],n2)
		if _,ok := neibour[n2] ; !ok {
			neibour[n2] = []int{}
		}
		neibour[n2] = append(neibour[n2],n1)
	}
	result := make([]int,len(cnt))
	for k,v := range cnt {
		if v == 1 {
			result[0] = k
			break
		}
	}
	result[1] = neibour[result[0]][0]
	for i := 2 ; i < len(cnt) ; i++ {
		for _,node := range neibour[result[i-1]] {
			if node != result[i-2] {
				result[i] = node
				break
			}
		}
	}
	return result
}
// @lc code=end

