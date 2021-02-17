/*
 * @lc app=leetcode.cn id=566 lang=golang
 *
 * [566] 重塑矩阵
 */

// @lc code=start
func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums) == 0 {
		return nums
	}
	_r,_c := len(nums),len(nums[0])
	if _r*_c != r*c {
		return nums
	}
	_i,_j := 0,0
	ans := [][]int{}

	for i := 0 ; i < r ; i++ {
		ans = append(ans,[]int{})
		for j := 0 ; j < c ; j++ {
			ans[i] = append(ans[i],nums[_i][_j])
			/*
			_j = (_j+1)%_c
			if _j == 0 {
				_i++
			}
			*/
			if _j < _c-1 {
				_j++
			}else{
				_i++
				_j = 0
			}
		}
	}
	return ans
}
// @lc code=end

