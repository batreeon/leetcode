/*
 * @lc app=leetcode.cn id=959 lang=golang
 *
 * [959] 由斜杠划分区域
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
	//小的部分作为子树
    uf.size[fx] += uf.size[fy]
    uf.parent[fy] = fx
	uf.setCount--
	//每成功进行一次合并，连通区域就减少一
}

func regionsBySlashes(grid []string) int {
    n := len(grid)//方形边长，n*n
    uf := newUnionFind(4 * n * n)//注意哟，n*n个小方块，每个小方块分成四片
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            idx := i*n + j	//标记第几个消防快
            if i < n-1 {//不是最底下一列
                bottom := idx + n
                uf.union(idx*4+2, bottom*4)//每一小方块，从上然后顺时针编号为0~3,2为下，0(4)为上
            }
            if j < n-1 {//不是最右边一列
                right := idx + 1
                uf.union(idx*4+1, right*4+3)//１为右，3为左
            }
            if grid[i][j] == '/' {
                uf.union(idx*4, idx*4+3)//上左
                uf.union(idx*4+1, idx*4+2)//右下
            } else if grid[i][j] == '\\' {
                uf.union(idx*4, idx*4+1)//上右
                uf.union(idx*4+2, idx*4+3)//下左
            } else {
                uf.union(idx*4, idx*4+1)
                uf.union(idx*4+1, idx*4+2)
                uf.union(idx*4+2, idx*4+3)
            }//秒啊，太秒了
        }
    }
    return uf.setCount//初始为4*n*n最终返回剩余总的连通区域个数
}
// @lc code=end

