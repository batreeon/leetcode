/*
 * @lc app=leetcode.cn id=89 lang=golang
 *
 * [89] 格雷编码
 */

// @lc code=start
func grayCode(n int) []int {
	// n=2 mask = 11 
	// 0
	// 11-1 & 11 = 10
	// 10-1 & 11 = 01
	// 01-1 = 00 结束
	// mask
	//不符合

	//二进制转格雷码，高位不变，其余位与高一位做异或
	// ans := []int{}
	// N := 1<<n
	// for i := 0 ; i < N ; i++ {
	// 	ans = append(ans,i ^ (i>>1))//不影响高位
	// }
	// return ans

	//官方解答：
	//0
	//0 1
	//00 01 11 10
	//000 001 011 010 110 111 101 100
	ans := []int{0}
	for i := 0 ; i < n ; i++ {
		head := 1 << i
		for j := len(ans)-1 ; j >= 0 ; j-- {
			ans = append(ans,ans[j]+head)
		}
	}
	return ans
}
// @lc code=end

