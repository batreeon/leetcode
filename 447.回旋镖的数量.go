/*
 * @lc app=leetcode.cn id=447 lang=golang
 *
 * [447] 回旋镖的数量
 */

// @lc code=start
func numberOfBoomerangs(points [][]int) int {
	/*
	n := len(points)
	dis := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		dis[i] = map[int]int{}
	}
	result := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j != i {
				p1,p2 := points[i],points[j]
				distance := (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
				dis[i][distance]++
			}
		}
		for _,d := range dis[i] {
			result += d*(d-1)
		}
	}

	return result
	*/
	n := len(points)
	result := 0
	for i := 0; i < n; i++ {
		dis := map[int]int{}
		for j := 0; j < n; j++ {
			if j != i {
				p1,p2 := points[i],points[j]
				distance := (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
				dis[distance]++
			}
		}
		for _,d := range dis {
			result += d*(d-1)
		}
	}

	return result
}
// @lc code=end

