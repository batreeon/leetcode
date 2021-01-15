/*
 * @lc app=leetcode.cn id=947 lang=golang
 *
 * [947] 移除最多的同行或同列石头
 */

// @lc code=start
func removeStones(stones [][]int) int {
	/*
	type point [2]int
	xp := [][]point{}
	yp := [][]point{}
	numrelated := map[[2]int]int{}
	minrelated := 0
	var nextdel point 
	for _,v := range stones {
		xp = append(xp[v[0]],v)
		yp = append(yp[v[1]],v)
	}
	for _,v := range stones {
		if len(xp[v[0]])-1+len(yp[v[1]])-1 != 0 {
			numrelated[v] = numrelated[v]+len(xp[v[0]])-1+len(yp[v[1]])-1
		}
	}
	for len(numrelated)!=0 {
		for k,v := range numrelated {
			if v > minrelated {

			}
		}
	}
	*/
	n := len(stones)
    graph := make([][]int, n)
    for i, p := range stones {
        for j, q := range stones {
            if p[0] == q[0] || p[1] == q[1] {
                graph[i] = append(graph[i], j)
            }
        }
    }
    vis := make([]bool, n)
    var dfs func(int)
    dfs = func(v int) {
        vis[v] = true
        for _, w := range graph[v] {
            if !vis[w] {
                dfs(w)
            }
        }
    }
    cnt := 0
    for i, v := range vis {
        if !v {
            cnt++
            dfs(i)
        }
    }
    return n - cnt
}
// @lc code=end

