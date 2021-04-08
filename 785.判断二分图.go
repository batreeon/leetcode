/*
 * @lc app=leetcode.cn id=785 lang=golang
 *
 * [785] 判断二分图
 */

// @lc code=start
/*
package main

type unionFind struct {
	parent []int
	rank   []int
	set    int
}

func newUnionFind(k int) *unionFind {
	parent := make([]int, k)
	rank := make([]int, k)
	for i := 0; i < k; i++ {
		parent[i] = i
		rank[i] = 0
	}
	return &unionFind{
		parent: parent,
		rank:   rank,
		set:    k,
	}
}

func (u *unionFind) find(k int) int {
	p := u.parent[k]
	if p != k {
		u.parent[k] = u.find(p)
	}
	return u.parent[k]
}

func (u *unionFind) union(x,y int) bool {
	px,py := u.find(x),u.find(y)
	if px == py {
		return false
	}
	if u.rank[px] < u.rank[py] {
		px,py = py,px
	}
	u.rank[px] += u.rank[py]
	u.parent[py] = px
	u.set--
	return true
}

func isBipartite(graph [][]int) bool {
	unionfind := newUnionFind(len(graph))
	for _,es := range graph {
		if len(es) == 0 {
			unionfind.set--
		}
		for i := 0 ; i < len(es)-1 ; i++ {
			unionfind.union(es[i],es[i+1])
		}
	}
	return unionfind.set == 2
}

// 大错特错，错误
*/

func isBipartite(graph [][]int) bool {
	// 广度优先遍历
	uncolored,red,green := 0,1,2
	n := len(graph)
	// 记录所有节点的标记颜色，初始未标记
	color := make([]int,n)
	for i := 0 ; i < n ; i++ {
		if color[i] == uncolored {
			q := []int{i}
			// 初始标为red
			color[i] = red
			for len(q) > 0 {
				node := q[0]
				nextColor := red
				// 相邻的颜色不一样
				if color[node] == red {
					nextColor = green
				}
				q = q[1:]
				// 遍历标记邻居
				for _,v := range graph[node] {
					if color[v] == uncolored {
						color[v] = nextColor
						q = append(q,v)
					// 若邻居颜色已有标记，但有冲突则错误返回
					}else if color[v] != nextColor {
						return false
					}
				}
			}
		}
	}
	return true
}
// @lc code=end
