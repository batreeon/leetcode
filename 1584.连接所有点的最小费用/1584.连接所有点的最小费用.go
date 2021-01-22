
/*
 * @lc app=leetcode.cn id=1584 lang=golang
 *
 * [1584] 连接所有点的最小费用
 */

// @lc code=start
//不会写啊
//官方Kruskal
// type unionFind struct {
//     parent, rank []int
// }

// func newUnionFind(n int) *unionFind {
//     parent := make([]int, n)
//     rank := make([]int, n)
//     for i := range parent {
//         parent[i] = i
//         rank[i] = 1
//     }
//     return &unionFind{parent, rank}
// }

// func (uf *unionFind) find(x int) int {
//     if uf.parent[x] != x {
//         uf.parent[x] = uf.find(uf.parent[x])
//     }
//     return uf.parent[x]
// }

// func (uf *unionFind) union(x, y int) bool {
//     fx, fy := uf.find(x), uf.find(y)
//     if fx == fy {
//         return false
//     }//判断是否在同一棵树中
//     if uf.rank[fx] < uf.rank[fy] {
//         fx, fy = fy, fx
//     }//是左右子树高度尽量均衡
//     uf.rank[fx] += uf.rank[fy]
//     uf.parent[fy] = fx
//     return true
// }

// func dist(p, q []int) int {
//     return abs(p[0]-q[0]) + abs(p[1]-q[1])
// }

// func minCostConnectPoints(points [][]int) (ans int) {
//     n := len(points)
//     type edge struct{ v, w, dis int }
//     edges := []edge{}
//     for i, p := range points {
//         for j := i + 1; j < n; j++ {
//             edges = append(edges, edge{i, j, dist(p, points[j])})
//         }
//     }//存放所有结点之间的费用

//     //按dist排序
//     sort.Slice(edges, func(i, j int) bool { return edges[i].dis < edges[j].dis })

//     //初始化一组unionFind
//     uf := newUnionFind(n)
//     left := n - 1//共有n个结点，只需要n-1条边即可
//     for _, e := range edges {
//         if uf.union(e.v, e.w) {//e.v和e.w分别为一条路径的两个端点
//             ans += e.dis
//             left--
//             if left == 0 {
//                 break
//             }
//         }
//     }
//     return
// }

// func abs(x int) int {
//     if x < 0 {
//         return -x
//     }
//     return x
// }

//自主尝试prim
//Kruskal是按权值，从小到大考察，如果路径两端点不在同一子树中就将其合并，
//加入n-1条边就完成了
//prim是首先随便选择一个结点作为起点，考察其邻居结点，将权值最小的路径合并进来，
//然后以新加入的结点来对数外结点的权重进行更新，然后继续选择横跨树外树内最小的边，满n-1条边就够了

// package main 

// import (
//     "fmt"
//     "math"
// )

func dist(p, q []int) int {
	return abs(p[0]-q[0]) + abs(p[1]-q[1])
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func minCostConnectPoints(points [][]int) (ans int) {
	n := len(points)
	var nextNode int//记录下一个距离树最近的节点
	lowcost := []int{}//记录结点到树的最小距离，已经在树中的置为0
	for i := 0; i < n; i++ {
		lowcost = append(lowcost, math.MaxInt64)
    }
	lowcost[0] = 0
    //选取第一个结点作为起点
	checkpoint := 0
	for i := 0; i < n-1; i++ {
		minDst := math.MaxInt64
		for j, l := range lowcost {
			if l > 0 { //out of tree
				lowcost[j] = min(lowcost[j], dist(points[j], points[checkpoint]))
				//每一次只需要根据上一个加入结点更新树外结点到树的最小距离
				//为什么只和checkpoint比，不是和全树比吗，因为每一次加入一个结点都比过了，所以前面老结点就没必要比了
				//fmt.Println(j," ",kp," ",dist(points[j], points[kp])," ",lowcost[j])
				if minDst > lowcost[j] {
					minDst = lowcost[j]
					nextNode = j
					//fmt.Println(-1," ",minDst," ",nextNode)
				}
			}
        }
        //fmt.Println(minDst," ",nextNode)
		lowcost[nextNode] = 0
		checkpoint = nextNode
		ans = ans + minDst
	}
	return
}
// func main() {
// 	p := [][]int{{0,0},{2,2},{3,10},{5,2},{7,0}}
//     fmt.Println(minCostConnectPoints(p))
// }
// @lc code=end

