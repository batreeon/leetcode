/*
 * @lc app=leetcode.cn id=802 lang=golang
 *
 * [802] 找到最终的安全状态
 */

// @lc code=start
func eventualSafeNodes(graph [][]int) []int {

	/*
	// 思路,没有出边的是安全的,有一条出边且指向安全点,那么本点也是安全的
	// 若一个点在其所有路径中存在环,那么这个点不是安全的
	// 若一个点的所有路径不存在环,且最后能到达安全点,那么该点是安全的
	n := len(graph)
	// result := []int{}
	safe := make([]bool,n)
	for i,v := range graph {
		if len(v) == 0 || (len(v) == 1 && len(graph[v[0]]) == 0){
			safe[i] = true
		}
	}

	// 判断点的所有路径是否有环,dfs
	ring := make([]bool,n)
	seen := make([]bool,n)
	var hasRing func(v int) bool 
	hasRing = func(v int) bool {
		seen[v] = true
		for _,next := range graph[v] {
			if seen[next] || ring[next] {
				ring[v] = true
				return true
			}
			if hasRing(next) {
				ring[v] = true
				return true
			}
		}
		seen[v] = false
		return false
	}

	for i := 0 ; i < n ; i++ {
		if !safe[i] && !ring[i] {
			hasRing(i)
		}
	}

	// 判断点是否能到达安全点
	var dfs func(v int) bool
	dfs = func(v int) bool {
		for _,next := range graph[v] {
			if safe[next] {
				safe[v] = true
				return true
			}
			if dfs(next) {
				safe[v] = true
				return true
			}
		}
		return false
	}

	for i := 0 ; i < n ; i++ {
		if !safe[i] && !ring[i] {
			hasRing(i)
		}
	}
	for i := 0 ; i < n ; i++ {
		if !safe[i] && !ring[i] {
			dfs(i)
		}
	}

	result := []int{}
	for i,issafe := range safe {
		if issafe {
			result = append(result,i)
		}
	}
	sort.Ints(result)
	return result
	*/

	/*
	n := len(graph)
	out := make([]int,n)
	pre := make([][]int,n)
	order := []int{}

	for i,vs := range graph {
		out[i] = len(vs)
		for _,v := range vs {
			pre[v] = append(pre[v],i)
		}
	}

	for i,outNum := range out {
		if outNum == 0 {
			order = append(order,i)
		}
	}

	for i := 0 ; i < len(order) ; i++ {
		del := order[i]
		for _,preNode := range pre[del] {
			out[preNode]--
			if out[preNode] == 0 {
				order = append(order,preNode)
			}
		}
	}
	sort.Ints(order)
	return order
	*/

	// 相当于进行了一个拓扑排序
	n := len(graph)
	// 记录每个点的出度
	out := make([]int,n)
	// 记录每个点的前导
	pre := make([][]int,n)
	// 记录已经安全的边，也就是出度为零的
	order := []int{}

	for i,nextV := range graph {
		out[i] = len(nextV)
		for _,v := range nextV {
			pre[v] = append(pre[v],i)
		}
	}
	for i,outNum := range out {
		if outNum == 0 {
			order = append(order,i)
		}
	}

	for i := 0 ; i < len(order) ; i++ {
		del := order[i]
		for _,preNode := range pre[del] {
			out[preNode]--
			if out[preNode] == 0 {
				order = append(order,preNode)
			}
		}
	}
	sort.Ints(order)
	return order
}
// @lc code=end

