/*
 * @lc app=leetcode.cn id=1871 lang=golang
 *
 * [1871] 跳跃游戏 VII
 */

// @lc code=start
func canReach(s string, minJump int, maxJump int) bool {
	n := len(s)
	if s[n-1] == '1' {
		return false
	}
	can := []int{n-1}
	// 从后往前考察
	// can数组要维护多个能到达终点的下标
	// 因为靠前的可能因为超过minJump而不一定用得上，但可能在以后用得上
	// 因此需要维护当前下标i + minJump之前的（也就是更小的），所有能到终点的下标值
	// 每次把多余的都删掉
	for i := n-2 ; i >= 0 ; i-- {
		if s[i] == '1' {
			continue
		}
		for j := len(can)-1 ; j >= 0 ; j-- {
			if can[j]-i >= minJump && can[j]-i <= maxJump && s[i] == '0' {
				if i == 0 {
					return true
				}
				// 已经达到minJump了，再大的就多余了
				can = append(can,i)
				can = can[j:]
				break
			}
		}
	}
	return false
}
// @lc code=end

