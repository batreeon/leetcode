/*
 * @lc app=leetcode.cn id=210 lang=golang
 *
 * [210] 课程表 II
 */

// @lc code=start
package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	// 错误，没有判断是否是有向无环图，只能处理是有向无环图的情况
	/*
	var order []int
	seen := map[int]bool{}
	m := make([][]int,numCourses)
	for i := 0 ; i < numCourses ; i++ {
		m[i] = []int{}
	}
	for _,pre := range prerequisites {
		m[pre[0]] = append(m[pre[0]],pre[1])
	}
	var visitedAll func([]int)
	visitedAll = func(pres []int) {
		for _,pre := range pres{
			if !seen[pre] {
				seen[pre] = true
				visitedAll(m[pre])
				order = append(order,pre)
			}
		}
	}
	var keys []int
	for i := 0 ; i < numCourses ; i++ {
		keys = append(keys,i)
	}
	sort.Ints(keys)
	visitedAll(keys)
	return order
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
	if len(order) != numCourses {
		return []int{}
	}
	return order
}
// @lc code=end

