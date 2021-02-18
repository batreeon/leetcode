/*
 * @lc app=leetcode.cn id=995 lang=golang
 *
 * [995] K 连续位的最小翻转次数
 */

// @lc code=start
func minKBitFlips(A []int, K int) int {
	//看得题解:
	//https://leetcode-cn.com/problems/minimum-number-of-k-consecutive-bit-flips/solution/hua-dong-chuang-kou-shi-ben-ti-zui-rong-z403l/
	l := len(A)
	ans := 0
	q := []int{}//若当前下标元素需要翻转则将下标入队
	for i , v := range A {
		if len(q) > 0 && q[0] + K <= i {
			q = q[1:]
		}//已经不在上一组范围内了，前k个已经考察完了
		if len(q) % 2 == v {//len（q)表示对当前下标元素翻转的次数，若满足该判断条件则需要翻转
			if i + K > l {
				return -1
			}//一次翻转需要连着后面共k个值翻转，对前免无影响，若后面元素不足，则返回－１
			q = append(q,i)//改下标元素需要翻转，入队
			ans++
		}
	}
	return ans
}
// @lc code=end

