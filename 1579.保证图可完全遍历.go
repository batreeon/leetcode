/*
 * @lc app=leetcode.cn id=1579 lang=golang
 *
 * [1579] 保证图可完全遍历
 */

// @lc code=start
type unionFind struct {
    parent, size []int
    setCount     int // 当前连通分量数目
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

func (uf *unionFind) inSameSet(x, y int) bool {
    return uf.find(x) == uf.find(y)
}

func maxNumEdgesToRemove(n int, edges [][]int) int {
    ans := len(edges)
    alice, bob := newUnionFind(n), newUnionFind(n)
    for _, e := range edges {
        x, y := e[1]-1, e[2]-1
        if e[0] == 3 && (!alice.inSameSet(x, y) || !bob.inSameSet(x, y)) {
            // 保留这条公共边
            alice.union(x, y)
            bob.union(x, y)
            ans--
        }
    }
    uf := [2]*unionFind{alice, bob}
    for _, e := range edges {
        if tp := e[0]; tp < 3 && uf[tp-1].union(e[1]-1, e[2]-1) {
            // 保留这条独占边
            ans--
        }
    }
    if alice.setCount > 1 || bob.setCount > 1 {
        return -1
    }
    return ans
    //算法思路：用ans记录总边数，没使用一条边就减去1，最后剩的就是可以删除的数量
    //分别为alice,bob建立两个并查集，首先考察类型3的边，
    //如果有一人需要（即这条边两端点不在同一连通分量内），则将两端点进行合并，并将ans减1
    //接着考察独占边，根据type值，分别对alice或者bob的并查集尝试进行合并，
    //若合并成功，则ans-1,合并失败则继续考察下一条独占边
}

// @lc code=end

