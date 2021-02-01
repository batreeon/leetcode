/*
 * @lc app=leetcode.cn id=839 lang=golang
 *
 * [839] 相似字符串组
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

func (uf *unionFind) union(x, y int) {
    fx, fy := uf.find(x), uf.find(y)
    if fx == fy {
        return
    }
    if uf.size[fx] < uf.size[fy] {
        fx, fy = fy, fx
    }
    uf.size[fx] += uf.size[fy]
    uf.parent[fy] = fx
    uf.setCount--
}

func (uf *unionFind) inSameSet(x, y int) bool {
    return uf.find(x) == uf.find(y)
}

func isSimilar(s, t string) bool {
    diff := 0
    for i := range s {
        if s[i] != t[i] {
            diff++
            if diff > 2 {
                return false
            }
        }
    }
    return true
}//判断是否相似，交换两个字母，但其他字母序号不会变，因此一一比对，若有两位以上不符，则不相似

func numSimilarGroups(strs []string) int {
    n := len(strs)
    uf := newUnionFind(n)
    for i, s := range strs {
        for j := i + 1; j < n; j++ {
            if !uf.inSameSet(i, j) && isSimilar(s, strs[j]) {
                uf.union(i, j)
            }
        }
    }
    return uf.setCount
}
// @lc code=end

