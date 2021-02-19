/*
 * @lc app=leetcode.cn id=697 lang=golang
 *
 * [697] 数组的度
 */

// @lc code=start
func findShortestSubArray(nums []int) int {
	type infor []int
	//三个元素，用来记录该值为首尾的字数组的左边界，右边界和出现次数

	m := map[int]infor{}
	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}
	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}

	n := len(nums)
	maxL := 0//度
	for i,k := range nums {
		if _,ok := m[k] ; !ok {
			if n-i >= maxL {
				m[k] = infor{i,i,1}
				maxL = max(maxL,m[k][2])
			}
			//后面长度不够，新出现的肯定不再会达到度了，也不用在加到map里了
		}else{
			m[k][1] = i
			m[k][2]++
			maxL = max(maxL,m[k][2])
		}
	}

	ans := math.MaxInt64
	for _,v := range m {
		if v[2] == maxL {
			ans = min(ans,v[1]-v[0]+1)
		}
	}
	return ans
}
// @lc code=end

