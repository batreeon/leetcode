/*
 * @lc app=leetcode.cn id=765 lang=golang
 *
 * [765] 情侣牵手
 */

// @lc code=start

//好难，做法不难，但完全没理解
type unionFind struct {
    parent, size []int
    setCount int // 当前连通分量数目
}

func newUnionFind(n int) *unionFind {
    parent := make([]int, n)
    size := make([]int, n)
    for i := range parent {
        parent[i] = i
        size[i] = 1
    }
    return &unionFind{parent, size, n}
}

func (uf *unionFind) find(x int) int {
    if uf.parent[x] != x {
        uf.parent[x] = uf.find(uf.parent[x])
    }
    return uf.parent[x]
}

func (uf *unionFind) union(x, y int) bool {
    fx, fy := uf.find(x), uf.find(y)
    if fx == fy {
        return false
    }
    if uf.size[fx] < uf.size[fy] {
        fx, fy = fy, fx
    }
    uf.size[fx] += uf.size[fy]
    uf.parent[fy] = fx
    uf.setCount--
    return true
}

func minSwapsCouples(row []int) int {
	n := len(row)
	uf := newUnionFind(n/2)
	for i := 0 ; i < n ; {
		uf.union(row[i]/2,row[i+1]/2)
		i += 2
	}
	return n/2 - uf.setCount
}
// @lc code=end

