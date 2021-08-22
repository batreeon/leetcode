/*
你在一个城市里，城市由 n 个路口组成，路口编号为 0 到 n - 1 ，某些路口之间有 双向 道路。输入保证你可以从任意路口出发到达其他任意路口，且任意两个路口之间最多有一条路。

给你一个整数 n 和二维整数数组 roads ，其中 roads[i] = [ui, vi, timei] 表示在路口 ui 和 vi 之间有一条需要花费 timei 时间才能通过的道路。你想知道花费 最少时间 从路口 0 出发到达路口 n - 1 的方案数。

请返回花费 最少时间 到达目的地的 路径数目 。由于答案可能很大，将结果对 109 + 7 取余 后返回。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-ways-to-arrive-at-destination
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
import (
    "math"
)
func countPaths(n int, roads [][]int) int {
    m := make([][]int,n)
    for i := 0 ; i < n ; i++ {
        m[i] = make([]int,n)
        for j := 0 ; j < n ; j++ {
            m[i][j] = math.MaxInt32
        }
    }
    
    V := make([][]int,n)
    for i := 0 ; i < n ; i++ {
        V[i] = make([]int,n)
        for j := 0 ; j < n ; j++ {
            V[i][j] = -1
        }
    }
    
    // 错误的
    // paths := make([][]int,n)
    // for i := 0 ; i < n ; i++ {
    //     paths[i] = make([]int,n)
    //     for j := 0 ; j < n ; j++ {
    //         paths[i][j] = 0
    //     }
    // }
    for _,road := range roads {
        u,v,time := road[0],road[1],road[2]
        m[u][v] = time
        m[v][u] = time
        // paths[u][v] = 1
        // paths[v][u] = 1
        V[u][v] = time
        V[v][u] = time
    }
    for i := 0 ; i < n ; i++ {
        for j := 0 ; j < n ; j++ {
            for k := 0 ; k < n ; k++ {
                if m[j][k] > m[j][i]+m[i][k] {
                    m[j][k] = m[j][i]+m[i][k]
                    // paths[j][k] = paths[j][i]*paths[i][k]
                }
                // else if m[j][k] == m[j][i]+m[i][k] {
                //     if paths[j][i]*paths[i][k] > paths[j][k] {
                //         paths[j][k] = paths[j][i]*paths[i][k]
                //     }
                // }
            }
        }
    }
    // return paths[0][n-1]
	
    path := m[0][n-1]
    result := 0
    var pathCount func(s,pathLen int)
    pathCount = func(s,pathLen int) {
        if V[s][n-1] == pathLen {
            result++
        }
        for i := 0 ; i < n-1 ; i++ {
            if i == s {
                continue
            }
            if V[s][i] == -1 {
                continue
            }
            if V[s][i] + m[i][n-1] == pathLen {
                pathCount(i,pathLen-V[s][i])
            }
        }
    }
    pathCount(0,path)
    return result
}
*/
package main

import (
    "math"
)
func countPaths(n int, roads [][]int) int {
    if n == 1 {
        return 1
    }
    m := make([][]int,n)
    for i := 0 ; i < n ; i++ {
        m[i] = make([]int,n)
        for j := 0 ; j < n ; j++ {
            m[i][j] = math.MaxInt64/2
            // 这个最大值取得不一样会影响结果
            // 如果取math.MaxInt32,那么有的样例可能出现路径长度超过这个值的，那就不行了，不能取这个
            // 但是你取math.MaxInt64连第一个样子都通过不了，没搞懂(7,[][]int{{0,6,7},{0,1,2},{1,2,3},{1,3,3},{6,3,3},{3,5,1},{6,5,1},{2,5,1},{0,4,5},{4,6,2}})
            /*下面更新图时会有这么一句：
            if m[j][k] > m[j][i]+m[i][k] {
                m[j][k] = m[j][i]+m[i][k]
            }
            如果你把不可达设置为math.MaxInt64,那么m[j][i]+m[i][k]就有可能是两个MaxInt64相加，
            因为是有符号数，可能会越界得到一个较小的数，导致错误的更新
            */
        }
    }
    
    // 记录邻接边
    V := make([][]int,n)
    for i := 0 ; i < n ; i++ {
        V[i] = make([]int,n)
        for j := 0 ; j < n ; j++ {
            V[i][j] = -1
        }
    }
    
    // 填充m和V
    for _,road := range roads {
        u,v,time := road[0],road[1],int(road[2])
        m[u][v] = time
        m[v][u] = time
        V[u][v] = time
        V[v][u] = time
    }
    // floyd
    for i := 0 ; i < n ; i++ {
        for j := 0 ; j < n ; j++ {
            for k := 0 ; k < n ; k++ {
                if m[j][k] > m[j][i]+m[i][k] {
                    m[j][k] = m[j][i]+m[i][k]
                }
            }
        }
    }
    // 0到n-1的最短路径
    minPath := m[0][n-1]

    // 为什么不用二维数组，minPath可能太大了，第二维空间可能不足了
    paths := make([]map[int]int,n)
    for i := 0 ; i < n ; i++ {
        paths[i] = map[int]int{}
    }

    // pathCount求点s到n-1路径为pathLen的路径数
    var pathCount func(s int,pathLen int) int
    pathCount = func(s int,pathLen int) int {
        // 有s到n-1直达边
        if V[s][n-1] == pathLen {
            paths[s][pathLen] += 1
        }
        // 遍历点s的邻接点
        for i := 0 ; i < n-1 ; i++ {
            if i == s {
                continue
            }
            if V[s][i] == -1 {
                continue
            }
            // s-i-(n-1)的路径长度为pathLen
            if V[s][i] + m[i][n-1] == pathLen {
                if path,ok := paths[i][pathLen-V[s][i]] ; ok {
                    // i到n-1的路径数已经计算过了，就直接读
                    paths[s][pathLen] += path
                }else{
                    paths[s][pathLen] += pathCount(i,pathLen-V[s][i])
                }
            }
        }
        paths[s][pathLen] %= 1000000007
        return paths[s][pathLen]
    }
    pathCount(0,minPath)
    return int(paths[0][minPath])
}