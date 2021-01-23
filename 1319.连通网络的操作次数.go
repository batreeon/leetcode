/*
 * @lc app=leetcode.cn id=1319 lang=golang
 *
 * [1319] 连通网络的操作次数
 */

// @lc code=start
type unionFind struct {
	parent, rank []int
}

func newUnionFind(n int) *unionFind {
	parent, rank := make([]int, n), make([]int, n)
	for i := range parent {
		parent[i], rank[i] = i, 1
	}
	return &unionFind{parent, rank}
}
func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}
func (uf *unionFind) union(x, y int) bool {
	px, py := uf.find(x), uf.find(y)
	if px == py {
		return false
	}
	if uf.rank[x] < uf.rank[y] {
		px, py = py, px
	}
	//小的作为子树
	uf.rank[px] += uf.rank[py]
	uf.parent[py] = px
	return true
}
func makeConnected(n int, connections [][]int) int {
	ln := len(connections) //边的个数
	var numTree int        //记录有多少棵不同的子树
	if n-1 > ln {
		return -1
	}
	uf := newUnionFind(n)
	for _, v := range connections {
		uf.union(v[0], v[1])
	}
	for i, _ := range uf.parent {
		if uf.parent[i] == i {
			numTree++
		}
	}
	return numTree - 1
}

// @lc code=end

