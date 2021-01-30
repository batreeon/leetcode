/*
 * @lc app=leetcode.cn id=1631 lang=golang
 *
 * [1631] 最小体力消耗路径
 */

// @lc code=start
type pair struct{ x, y int }
var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func minimumEffortPath(heights [][]int) int {
    n, m := len(heights), len(heights[0])
    return sort.Search(1e6, func(maxHeightDiff int) bool {
        vis := make([][]bool, n)
        for i := range vis {
            vis[i] = make([]bool, m)
        }
        vis[0][0] = true
        queue := []pair{{}}//该切片的元素是点结构体，默认初始化一个值{0,0}原点吧
        for len(queue) > 0 {
            p := queue[0]
            queue = queue[1:]
            if p.x == n-1 && p.y == m-1 {
                return true
            }//当到达右下点
            for _, d := range dirs {
                x, y := p.x+d.x, p.y+d.y
                if 0 <= x && x < n && 0 <= y && y < m && !vis[x][y] && abs(heights[x][y]-heights[p.x][p.y]) <= maxHeightDiff {
                    vis[x][y] = true
                    queue = append(queue, pair{x, y})
                }
            }
        }
        return false
    })
}
//这道题完美利用了golang中sort库的sort.Search方法，
//该方法返回可使第二个函数参数f(x)返回true的最小的x值，若没有返回true的，则返回第一个参数值
//https://www.bookstack.cn/read/The-Golang-Standard-Library-by-Example/chapter03-03.1.md

//还有一种解法，将所有边按大小排序，从小到大合并，当合并某一条边后，
//[0,0]点和[n-1,m-1]点父亲相同时，即属于同一连通子图，该条边的权重即为答案
func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
// @lc code=end

