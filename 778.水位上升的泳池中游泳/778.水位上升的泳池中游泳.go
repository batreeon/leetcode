/*
 * @lc app=leetcode.cn id=778 lang=golang
 *
 * [778] 水位上升的泳池中游泳
 */

// @lc code=start
// package main 
// import (
// 	"fmt"
// 	"sort"
// )
type unionFind struct {
	parent,rank []int
}
func newUnionFind(n int) *unionFind {
	parent,rank := make([]int,n),make([]int,n)
	for i,_ := range parent {
		parent[i] = i 
		rank[i] = 1
	}
	return &unionFind{parent,rank}
}
func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {//我一开始把if写成for了，若判断失败则会进入无限循环
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}
func (uf *unionFind) union(x,y int) {
	fx,fy := uf.find(x),uf.find(y)
	//fmt.Println(fx," ",fy)
	if fx == fy {
		return 
	}
	if uf.rank[fx] < uf.rank[fy] {
		fx,fy = fy,fx
	}//小的作为子树
	uf.rank[fx] += uf.rank[fy]
	uf.parent[fy] = fx
	return 
}
func (uf *unionFind) linked(x,y int) bool {
	return uf.find(x) == uf.find(y)
}//判断是否连通

type edge struct {
	u,v,w int
}//边：起点，终点，权值。权值即为两点中平台高度的较大值
func swimInWater(grid [][]int) int {
	n,m := len(grid),len(grid[0])
	//fmt.Println(n," ",m)
	edges := []edge{}
	for i,a := range grid {
		for j,b := range a {
			d := i*m+j	//边终点
			if i > 0 {
				edges = append(edges,edge{d-m,d,max(grid[i-1][j],b)})
			}
			if j > 0 {
				edges = append(edges,edge{d-1,d,max(grid[i][j-1],b)})
			} 
		}
	}
	sort.Slice(edges,func(i,j int) bool {return edges[i].w < edges[j].w})
	//fmt.Println(edges)
	uf := newUnionFind(n*m)
	for _,e := range edges {
		//fmt.Println(0," ",e.u," ",e.v)
		uf.union(e.u,e.v)
		//fmt.Println(1," ",e.u," ",e.v)
		if uf.linked(0,n*m-1) {
			//fmt.Println(3)
			return e.w
		}
		//fmt.Println(2)
	}
	return 0
}
func max(x,y int) int {
	if x > y {
		return x
	}else{
		return y
	}
}
// func main() {
// 	g := make([][]int,2)
// 	g[0] = append(g[0],0)
// 	g[0] = append(g[0],2)
// 	g[1] = append(g[1],1)
// 	g[1] = append(g[1],3)
// 	fmt.Println(swimInWater(g))
// }
// @lc code=end

