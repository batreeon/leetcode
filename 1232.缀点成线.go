import "math"

/*
 * @lc app=leetcode.cn id=1232 lang=golang
 *
 * [1232] 缀点成线
 */

// @lc code=start
//计算两点斜率，x坐标相同则返回最大值
func caculate_a(p1, p2 []int) float64 {
	_y := float64(p1[1]) - float64(p2[1])
	_x := float64(p1[0]) - float64(p2[0])
	if _x == 0 {
		return math.MaxFloat64
	}
	return _y / _x
}

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) < 3 {
		return true
	}
	//y=ax+b
	a := caculate_a(coordinates[0], coordinates[1])
	for i := 2; i < len(coordinates); i++ {
		if caculate_a(coordinates[i], coordinates[0]) != a {
			return false
		}
	}
	return true
}

// @lc code=end

