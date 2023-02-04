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
	dfsred = func(start,end,distance,index int) {
		if red[start][end] {
			if result[end][index] == -1 || distance < result[end][index] {
				result[end][index] = distance
			}
			return
		}
		for next := range red[start] {
			if !visitedblue[next] {
				visitedblue[next] = true
				dfsblue(next,end,distance+1,index)
				delete(visitedblue,next)
			}
		}
		return
	}
	dfsblue = func(start,end,distance,index int) {
		if blue[start][end] {
			if result[end][index] == -1 || distance < result[end][index] {
				result[end][index] = distance
			}
			return
		}
		for next := range blue[start] {
			if !visitedred[next] {
				visitedred[next] = true
				dfsred(next,end,distance+1,index)
				delete(visitedred,next)
			}
		}
		return
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

	/*
	// 20230203有错误，有环的路径没考虑到
	red := make([][]int,n)
	blue := make([][]int,n)
	for _,edge := range red_edges {
		if red[edge[0]] == nil {
			red[edge[0]] = []int{}
		}
		red[edge[0]] = append(red[edge[0]], edge[1])
	}
	for _,edge := range blue_edges {
		if blue[edge[0]] == nil {
			blue[edge[0]] = []int{}
		}
		blue[edge[0]] = append(blue[edge[0]], edge[1])
	}

	// distance[i]分别记录0到i的路径长度，一个表示最后一条边是red，另一个blue
	distance := make([][2]int, n)
	// 从下标1开始，因为0到0距离为0
	for i := 1; i < n; i++ {
		distance[i][0] = -1
		distance[i][1] = -1
	}
	// bfs
	nodes := []int{0}
	visitedr := map[int]bool{0:true}
	visitedb := map[int]bool{0:true}
	for len(nodes) > 0 {
		l := len(nodes)
		for _, node := range nodes {
			for rnode := range red[node] {
				if distance[node][1] != -1 && (distance[rnode][0] == -1 || distance[node][1] + 1 < distance[rnode][0]) {
					distance[rnode][0] = distance[node][1] + 1
				}
				if !visitedr[rnode] {
					visitedr[node] = true
					nodes = append(nodes, rnode)
				}
			}
			for bnode := range blue[node] {
				if distance[node][0] != -1 && (distance[bnode][1] == -1 || distance[node][0] + 1 < distance[bnode][1]) {
					distance[bnode][1] = distance[node][0] + 1
				}
				if !visitedb[bnode] {
					visitedb[bnode] = true
					if nodes[len(nodes)-1] != bnode {
						nodes = append(nodes, bnode)
					}
				}
			}
		}
		nodes = nodes[l:]
	}

	result := make([]int, n)
	for i := 1; i < n; i++ {
		result[i] = func(x, y int) int {
			if x == -1 {
				return y
			}
			if y == -1 {
				return x
			}
			if x < y {
				return x
			}
			return y
		}(distance[i][0], distance[i][1])
	}

	return result
	*/


	// 记录以每个点为起点时的所有终点
	red := make([][]int,n)
	blue := make([][]int,n)
	for _,edge := range red_edges {
		if edge[1] == 0 {
			continue
		}
		if red[edge[0]] == nil {
			red[edge[0]] = []int{}
		}
		red[edge[0]] = append(red[edge[0]], edge[1])
	}
	for _,edge := range blue_edges {
		if edge[1] == 0 {
			continue
		}
		if blue[edge[0]] == nil {
			blue[edge[0]] = []int{}
		}
		blue[edge[0]] = append(blue[edge[0]], edge[1])
	}

	// distance[i]分别记录0到i的路径长度，一个表示最后一条边是red，另一个blue
	distanceRed := make([]int, n)
	distanceBlue := make([]int, n)
	distanceRed[0] = 0
	distanceBlue[0] = 0
	// 从下标1开始，因为0到0距离为0
	for i := 1; i < n; i++ {
		distanceRed[i] = -1
		distanceBlue[i] = -1
	}

	// 记录点i是否被red或blue边作为终点访问过
	var visitedr map[int]bool
	var visitedb map[int]bool

	// bfs
	var bfsred func (nodes []int)
	var bfsblue func (nodes []int)

	// 下一条边是red
	bfsred = func (nodes []int) {
		l := len(nodes)
		if l == 0 {
			return
		}
		for _, node := range nodes {
			for _, next := range red[node] {
				if distanceBlue[node] != -1 && (distanceRed[next] == -1 || distanceBlue[node] + 1 < distanceRed[next]) {
					distanceRed[next] = distanceBlue[node] + 1
				}
				if !visitedr[next] {
					visitedr[next] = true
					nodes = append(nodes, next)
				}
			}
		}
		nodes = nodes[l:]
		bfsblue(nodes)
	}

	// 下一条边是blue
	bfsblue = func (nodes []int) {
		l := len(nodes)
		if l == 0 {
			return
		}
		for _, node := range nodes {
			for _, next := range blue[node] {
				if distanceRed[node] != -1 && (distanceBlue[next] == -1 || distanceRed[node] + 1 < distanceBlue[next]) {
					distanceBlue[next] = distanceRed[node] + 1
				}
				if !visitedb[next] {
					visitedb[next] = true
					nodes = append(nodes, next)
				}
			}
		}
		nodes = nodes[l:]
		bfsred(nodes)
	}


	// 第一步先走红边
	visitedr = map[int]bool{0:true}
	visitedb = map[int]bool{0:true}
	bfsred([]int{0})

	// 第一步先走蓝边
	// 为什么还要再重置一下这个呢。因为上面bfsred运行后，部分节点访问过了，但可能不是最优的
	// 如果不重置，bfsblue运行时这些节点将不会再append到nodes中，有可能会得不到最优结果
	visitedr = map[int]bool{0:true}
	visitedb = map[int]bool{0:true}
	bfsblue([]int{0})

	result := make([]int, n)
	for i := 1; i < n; i++ {
		result[i] = func(x, y int) int {
			if x == -1 {
				return y
			}
			if y == -1 {
				return x
			}
			if x < y {
				return x
			}
			return y
		}(distanceRed[i], distanceBlue[i])
	}

	return result
}

// @lc code=end

