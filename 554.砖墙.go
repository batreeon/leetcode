/*
 * @lc app=leetcode.cn id=554 lang=golang
 *
 * [554] 砖墙
 */

// @lc code=start
func leastBricks(wall [][]int) int {
	/*有一个问题，再分配gaps数组时，可能因为每一层过宽而需要太大的内存
	// 我才去的策略是记录每一层缝的位置
	width := 0
	for _,brick := range wall[0] {
		width += brick
	}
	height := len(wall)
	if width == 1 {
		return height
	}
	gaps := make([]int,width-1)
	maxgap := 0
	for _,layer := range wall {
		sum := 0
		for _,brick := range layer {
			sum += brick
			if sum == width {
				break
			}
			gaps[sum-1]++
			if gaps[sum-1] > maxgap {
				maxgap = gaps[sum-1]
			}
		} 
	}
	return height-maxgap
	*/

	// 最终submit时，才知道答案可以为0，也就是可以不穿过砖
	// 就是把上面的gaps数据结构由切片改成了map
	maxgap := 0
	gaps := map[int]int{}
	height := len(wall)
	width := 0
	for _,brick := range wall[0] {
		width += brick
	}
	for _,layer := range wall {
		sum := 0
		for _,brick := range layer {
			sum += brick
			if sum == width {
				// 到了最后的末尾就不用再记了
				break
			}
			gaps[sum-1] = gaps[sum-1]+1
			if gaps[sum-1] > maxgap {
				maxgap = gaps[sum-1]
			}
		}
	}
	return height-maxgap
}
// @lc code=end

