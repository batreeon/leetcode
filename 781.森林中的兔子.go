/*
 * @lc app=leetcode.cn id=781 lang=golang
 *
 * [781] 森林中的兔子
 */

// @lc code=start
func numRabbits(answers []int) int {
	cnt := map[int]int{}
	for _,answer := range answers {
		cnt[answer]++
	}
	result := 0
	for k,v := range cnt {
		if k == 0 {
			// 若回答为0,则这一种颜色只有一个,有几个回答,就有几种颜色
			result += v
		}else{
			// 为了使兔子数最小,让这些回答大于0的尽可能位同一种颜色
			// 回答为k,那么相应的颜色就有k+1只兔子
			// 那么v个兔子就可能:
			// 若v%(k+1) == 0 , 就至少要v/(k+1)组
			// 若v%(k+1) != 0 , 就至少要v/(k+1)+1组
			// 每一组有k+1只兔子
			
			// 算出有多少组
			group := 0
			group += v/(k+1)
			if v%(k+1) != 0 {
				group++
			}
			result += group*(k+1)
		}
	}
	return result
}
// @lc code=end

