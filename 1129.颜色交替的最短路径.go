/*
 * @lc app=leetcode.cn id=1129 lang=golang
 *
 * [1129] 颜色交替的最短路径
 */

// @lc code=start
func shortestAlternatingPaths(n int, red_edges [][]int, blue_edges [][]int) []int {
	
	/*
	// 啊啊啊超时了
	red := make([]map[int]bool,n)
	blue := make([]map[int]bool,n)
	for _,edge := range red_edges {
		if red[edge[0]] == nil {
			red[edge[0]] = map[int]bool{}
		}
		red[edge[0]][edge[1]] = true
	}
	for _,edge := range blue_edges {
		if blue[edge[0]] == nil {
			blue[edge[0]] = map[int]bool{}
		}
		blue[edge[0]][edge[1]] = true
	}

	result := make([][2]int,n)
	var visitedred map[int]bool
	var visitedblue map[int]bool
	var dfsred func(int,int,int,int) bool
	var dfsblue func(int,int,int,int) bool
	dfsred = func(start,end,distance,index int) bool {
		if red[start][end] {
			if result[end][index] == -1 || distance < result[end][index] {
				result[end][index] = distance
			}
			return true
		}
		for next := range red[start] {
			if !visitedblue[next] {
				visitedblue[next] = true
				dfsblue(next,end,distance+1,index)
				delete(visitedblue,next)
			}
		}
		return false
	}
	dfsblue = func(start,end,distance,index int) bool {
		if blue[start][end] {
			if result[end][index] == -1 || distance < result[end][index] {
				result[end][index] = distance
			}
			return true
		}
		for next := range blue[start] {
			if !visitedred[next] {
				visitedred[next] = true
				dfsred(next,end,distance+1,index)
				delete(visitedred,next)
			}
		}
		return false
	}

	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}
	res := make([]int,n)
	for i := 1 ; i < n ; i++ {
		visitedred = map[int]bool{}
		visitedblue = map[int]bool{}
		result[i] = [2]int{-1,-1}
		visitedred[0] = true
		dfsred(0,i,1,0)
		visitedblue[0] = true
		dfsblue(0,i,1,1)
		res[i] = -1
		if result[i][0] != -1 && result[i][1] != -1 {
			res[i] = min(result[i][0],result[i][1])
		}else if result[i][0] != -1 {
			res[i] = result[i][0]
		}else if result[i][1] != -1 {
			res[i] = result[i][1]
		}
	}
	return res
	*/
	
}
// @lc code=end

