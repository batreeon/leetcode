package dijkstra

/*
n:节点个数,(节点分别用0~n-1表示)
vs:边集合，每个元素为三元组[起点，终点，权值]
src：为起点
sign: true表示有向图，false表示无向图
算法返回起点src到各点的距离，不可达记为math.MaxInt64/2
*/
import (
	"math"
)
func dijkstra(n int,vs [][]int,src int,sign bool) []int {
	cost := make([][]int,n)
	for i := 0 ; i < n ; i++ {
		cost[i] = make([]int,n)
		for j := 0 ; j < n ; j++ {
			if i == j {
				cost[i][j] = 0
			}else{
				cost[i][j] = math.MaxInt64/2
			}
		}
	}
	if sign {
		for _,v := range vs {
			cost[v[0]][v[1]] = v[2]
		}
	}else{
		for _,v := range vs {
			cost[v[0]][v[1]] = v[2]
			cost[v[1]][v[0]] = v[2]
		}
	}
	
	distance := make([]int,n)
	distance[src] = 0
	used := make([]bool,n)

	for {
		v := -1
		for u := 0 ; u < n ; u++ {
			if !used[u] && (v == -1 || distance[u] < distance[v]) {
				v = u
			}
		}

		// 没有可更新的了
		if v == -1 {break}
		used[v] = true

		for u := 0 ; u < n ; u++ {
			distance[u] = min(distance[u],distance[v]+cost[v][u])
		}
	}

	return distance
}

func min(x,y int) int {
	if x < y {
		return x
	}
	return y
}