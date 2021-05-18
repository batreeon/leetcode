/*
 * @lc app=leetcode.cn id=421 lang=golang
 *
 * [421] 数组中两个数的最大异或值
 */

// @lc code=start
func findMaximumXOR(nums []int) int {
	const M int = 3000010
	son := make([][2]int,M)
	var idx int

	insert := func(x int) {
		p := 0
		for i := 30 ; i >= 0 ; i-- {
			s := &son[p][x >> i & 1]
			if *s == 0 {
				idx++
				*s = idx
			}
			p = *s
		}
	}
	query := func(x int) int {
		res := 0
		p := 0
		for i := 30 ; i >= 0 ; i-- {
			s := x >> i & 1
			sidx := 0
			if s == 0 {
				sidx = 1
			}
			if son[p][sidx] != 0 {
				res += 1 << i
				p = son[p][sidx]
			}else{
				p = son[p][s]
			}
		}
		return res
	}
	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}

	for i := range nums {
		insert(nums[i])
	}
	res := 0
	for i := range nums {
		res = max(res,query(nums[i]))
	}
	return res
}
// @lc code=end

