/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */

// @lc code=start
package main
/*
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
*/

func canFinish(numCourses int, prerequisites [][]int) bool {
	/*
	uf := newUnionFind(numCourses)
	for _,pre := range prerequisites {
		if pre[0] == uf.find(pre[1]) {
			return false
		}
		uf.union(pre[0],pre[1])
	}
	return true
	*/
	//拓扑排序结果
	//后面有append,如果不指明初始长度，后面append，长度就大于numCourses了
	order := make([]int,0,numCourses)	
	//记录入度
	in := make([]int,numCourses)	
	//记录后继课程，若一个入度为0的顶点被删了，那么他的后继的入度都应该减一
	next := make([][]int,numCourses)	
	for _,pre := range prerequisites {
		in[pre[0]]++
		next[pre[1]] = append(next[pre[1]],pre[0])
	}
	for i,inV := range in {
		if inV == 0 {
			order = append(order,i)
		}
	}
	for i := 0 ; i < len(order) ; i++ {
		del := order[i]	//被删除的顶点
		// 下面依次将被删除的节点的后继的入度减一
		nextNodes := next[del]
		for _,node := range nextNodes {
			in[node]--
			if in[node] == 0 {
				order = append(order,node)
			}
		}
	}
	return len(order) == numCourses
}
// @lc code=end

