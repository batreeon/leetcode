/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */

// @lc code=start
package main
type unionFind struct {
	parent,rank []int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int,n)
	rank := make([]int,n)
	for i := 0 ; i < n ; i++ {
		parent[i],rank[i] = i,1
	}
	return &unionFind{parent,rank}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind) union(x,y int) {
	uf.parent[x] = uf.find(y)
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	uf := newUnionFind(numCourses)
	for _,pre := range prerequisites {
		if pre[0] == uf.find(pre[1]) {
			return false
		}
		uf.union(pre[0],pre[1])
	}
	return true
}
// @lc code=end

