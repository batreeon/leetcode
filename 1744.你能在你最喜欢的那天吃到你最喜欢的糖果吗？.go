/*
 * @lc app=leetcode.cn id=1744 lang=golang
 *
 * [1744] 你能在你最喜欢的那天吃到你最喜欢的糖果吗？
 */

// @lc code=start
func canEat(candiesCount []int, queries [][]int) []bool {
	sum := make([]int,len(candiesCount))
	sum[0] = candiesCount[0]
	for i := 1 ; i < len(candiesCount) ; i++ {
		sum[i] = sum[i-1]+candiesCount[i]
	}
	result := make([]bool,len(queries))
	for i,query := range queries {
		// 在第query[1]天吃到query[0]，那么在query[1]当天或之前能够将query[1]-1的糖果全吃完，并且第query[1]当天依然还有吃的容量
		if query[0] == 0 {
			if sum[query[0]] <= query[1] {
				// 没到第query[1]天，就把0类糖吃完了
				continue
			}
		}else{
			if sum[query[0]-1] >= (query[1]+1)*query[2] {
				// 到第query[1]天都没把上一类吃完,那肯定吃不到第query[0]类了
				continue
			}
			if sum[query[0]] <= query[1] {
				// 没到第query[1]天呢，就把query[0]类糖吃完了
				continue
			}
		}
		result[i] = true
	}
	return result
}
// @lc code=end

