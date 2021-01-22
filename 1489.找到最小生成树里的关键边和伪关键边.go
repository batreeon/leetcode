/*
 * @lc app=leetcode.cn id=1489 lang=golang
 *
 * [1489] 找到最小生成树里的关键边和伪关键边
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
    }//祖先相同，说明属于同一棵子树，无法合并
    if uf.size[fx] < uf.size[fy] {
        fx, fy = fy, fx
    }
    uf.size[fx] += uf.size[fy]
    uf.parent[fy] = fx
    uf.setCount--
    return true
}

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
    for i, e := range edges {
        edges[i] = append(e, i)
	}//e是[3]int,from,to,weight，记录每条边的下标

	//按weight排序
    sort.Slice(edges, func(i, j int) bool { return edges[i][2] < edges[j][2] })

    calcMST := func(uf *unionFind, ignoreID int) (mstValue int) {//ignoreID表示不使用这条边
        for i, e := range edges {
            if i != ignoreID && uf.union(e[0], e[1]) {
                mstValue += e[2]
            }
        }
        if uf.setCount > 1 {//不连通
            return math.MaxInt64
        }
        return
    }

    mstValue := calcMST(newUnionFind(n), -1)//最小生成树的权值

    var keyEdges, pseudokeyEdges []int
    for i, e := range edges {
        // 是否为关键边
        if calcMST(newUnionFind(n), i) > mstValue {
			//去除掉这条边后，不连通或者连通但权值变大，则这条边必然出现在所有生成树中，是关键边
            keyEdges = append(keyEdges, e[3])
            continue//如果是关键边就不执行下面的部分了，下面的是判断是否是伪关键边
        }

		//首先，能运行到下面的代码处，已经在上面确定了该条边不是关键边，
		//若将它作为最小生成树中的一条边，最终也能生成一棵最小生成树，
		//则其出现在一棵最小生成树中，但没有在所有最小生成树中
        // 是否为伪关键边
        uf := newUnionFind(n)
        uf.union(e[0], e[1])
        if e[2]+calcMST(uf, i) == mstValue {
            pseudokeyEdges = append(pseudokeyEdges, e[3])
        }
    }

    return [][]int{keyEdges, pseudokeyEdges}
}
// @lc code=end

