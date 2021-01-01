/*
 * @lc app=leetcode.cn id=973 lang=golang
 *
 * [973] 最接近原点的 K 个点
 */

// @lc code=start
type sortedpoints [][]int

func (p sortedpoints) Len() int {
	return len(p)
}
func (p sortedpoints) Swap(i,j int) {
	p[i],p[j] = p[j],p[i]
}
func (p sortedpoints) Less(i,j int) bool {
	return (math.Pow(float64(p[i][0]),2)+math.Pow(float64(p[i][1]),2)) < (math.Pow(float64(p[j][0]),2)+math.Pow(float64(p[j][1]),2))
}

func kClosest(points [][]int, K int) [][]int {
	sort.Sort(sortedpoints(points))
	return points[:K]
}
// @lc code=end

