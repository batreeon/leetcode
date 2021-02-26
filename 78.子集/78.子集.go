/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
package main
func subsets(nums []int) [][]int {
	ans := [][]int{}
	ans = append(ans,[]int{})
	for _,v1 := range nums {
		for _,v2 := range ans {
			/*错误写法
			res := append(v2,v1)
			ans = append(ans,res)
			*/
			n := len(v2)+1
			res := make([]int,n)
			copy(res[:n-1],v2[:])
			res[n-1] = v1	
			ans = append(ans,res)
		}
	}
	return ans
}
// @lc code=end

