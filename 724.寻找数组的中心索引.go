/*
 * @lc app=leetcode.cn id=724 lang=golang
 *
 * [724] 寻找数组的中心索引
 */

// @lc code=start
// 把总和计算出来会不会溢出啊，换个方法 //更新，貌似可用
func sum(nums []int) (ans int) {
	for _,v := range nums {
		ans += v
	}
	return 
}
//1 7 3 6 -3 3 5 6  这就有两个中心索引
func pivotIndex(nums []int) int {

	//下面这个方法不行，因为元素有负数，如果全部同号还可以试试
	// li,ri := 0,len(nums)-1
	// //left,right := nums[li],nums[ri]
	// //minus := left-right
	// minus := nums[li] - nums[ri]
	// for ;; {
	// 	if li > ri {
	// 		break
	// 	}
	// 	if minus < 0 {
	// 		li++
	// 	}else if minus > 0 {
	// 		ri++
	// 	}
	// }
	
	total := sum(nums)
	left := 0
	for i,v := range nums {
		if left*2 + v == total {
			return i
		} 
		left += v
	}
	return -1
}
// @lc code=end

