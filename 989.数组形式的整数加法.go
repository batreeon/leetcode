/*
 * @lc app=leetcode.cn id=989 lang=golang
 *
 * [989] 数组形式的整数加法
 */

// @lc code=start
func addToArrayForm(A []int, K int) (ans []int) {
	//将Ｋ值从Ａ的低位开始加起，计算低位的值，在整除以10,再与高一位相加，直到Ｋ为0并且遍历完Ａ最高位
	for i := len(A)-1;i >= 0 || K > 0;i--{
		if i >= 0 {
			K += A[i]
		}
		ans = append(ans,K%10)
		K /= 10
	}
	reverse(ans)
	return
}

func reverse(a []int) {
	for i,n := 0,len(a)-1 ; i < n-i ; i++ {
		a[i],a[n-i] = a[n-i],a[i]
	}
}
// @lc code=end

