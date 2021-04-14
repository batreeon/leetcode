/*
 * @lc app=leetcode.cn id=1779 lang=golang
 *
 * [1779] 找到最近的有相同 X 或 Y 坐标的点
 */

// @lc code=start
func nearestValidPoint(x int, y int, points [][]int) int {
	abs := func(x int) int {
		if x < 0 {
			return -1*x
		}
		return x
	}
	isvalid := func(pos *[]int) (bool,int) {
		xx,yy := (*pos)[0],(*pos)[1]
		if x == xx || y == yy {
			return true,abs(x-xx)+abs(y-yy)
		}
		return false,-1
	}
	minDistance := math.MaxInt64
	minIndex := -1
	for i,pos := range points {
		if ok,dis := isvalid(&pos) ; ok && dis < minDistance {
			minDistance = dis
			minIndex = i
		}
	}
	return minIndex
}
// @lc code=end

