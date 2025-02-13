/*
 * @lc app=leetcode.cn id=1742 lang=golang
 *
 * [1742] 盒子中小球的最大数量
 */

// @lc code=start
func countBalls(lowLimit int, highLimit int) int {
    boxes := make([]int, 46)
	for i := lowLimit; i <= highLimit; i++ {
		s := 0
		for i := i; i > 0; i/=10 {
			s += i%10
		}
		boxes[s]++
	}
	sort.Ints(boxes)
	return boxes[45]
}
// @lc code=end

