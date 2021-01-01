/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
// package main

// import "fmt"
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
func nextPermutation(nums []int) {
	i := len(nums) - 1
	for ; i > 0 && nums[i] <= nums[i-1]; i-- {
	}

	if i != 0 {
		j := i //初始值肯定有nums[j]>nums[i]
		//下面循环是为了找出i右侧大于nums[i]的最小值
		for ; j < len(nums)-1 && nums[j+1] > nums[i-1]; j++ {
		}
		LEN := len(nums)
		reverse(nums[j+1:])//这一部分都是不大于nums[i-1]的
		reverse(nums[i:j])//这一部分都是大于nums[i-1]的
		//上面的反转使这两部分都成了升序，接下来移动后就能拼成一个连续的升序，
		//构成一个以nums[j]为最高位的最小数
		tail := make([]int,LEN-j,LEN-j)
		copy(tail,nums[j:])
		// fmt.Println(i, j,tail,nums)
		copy(nums[i-1+LEN-j:], nums[i-1:j])
		// fmt.Println(i, j,tail,nums)
		copy(nums[i-1:i-1+LEN-j], tail)
		// fmt.Println(i, j,tail,nums)
	} else {
		reverse(nums)
	}
}

//解题过程：
//递减情况比较简单不谈了。
//非递减情况:从尾到头找，找到第一个nums[i] > nums[i-1]的情况，
//i后面的都是递减的，因此以nums[i-1]为最高位的数字已经是最大的了，
//为了使数字增大，必须找到一个比nums[i-1]大的数来做最高位
//因此从i开始向后遍历，直到找到比nums[i-1]大的最小的数（因为题目要求的是下一个更大的序列），将这个数下标记为j
//替换nums[i-1]和nums[j]，接着把开头右边的数降序排列。得到的数就是比原数大的最小的数
//
// func main() {
// 	s := []int{1,3,2}
// 	nextPermutation(s)
// 	fmt.Println(s)
// }

// @lc code=end
