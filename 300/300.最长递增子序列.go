/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start
package main

import "sort"
func lengthOfLIS(nums []int) int {
	//动规
	/*
	l := len(nums)
	f := make([]int,l)
	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}
	ans := 0
	for i := 0 ; i < l ; i++ {
		maxf := 0
		for j := 0 ; j < i ; j++ {
			if nums[j] < nums[i] {
				maxf = max(maxf,f[j])
			}
		}
		f[i] = maxf + 1
		ans = max(ans,f[i])
	}
	//f[i] = max(f[j])+1,j < i 并且 nums[j]<nums[j]
	return ans
	*/

	/*
	//贪心加二分查找
	//维护一个递增序列，使这个递增序列的最大值，是同长度递增序列中最小的，这就是贪心体现的地方
	l := len(nums)
	if l == 0 {
		return 0
	}
	ans := 1
	d := make([]int,0,l)
	d = append(d,nums[0])
	for _,v := range nums {
		if v > d[ans-1] {
			// 递增序列变长了，直接把这个新的更大的元素加到末尾
			d = append(d,v)
			ans++
		}else{
			// 为了贪心，递增序列中的元素每次尽可能增得慢一点
			// 用较小的元素替换已有的递增序列
			pos := -1
			l,r := 0,ans-1

			// 二分查找找到从后往前数第一个比当前元素小的元素，然后插到他后面，替换原来的较大的递增值
			for l <= r {
				mid := (l+r)>>1
				if d[mid] < v {
					pos = mid
					l = mid+1
				}else{
					r = mid-1
				}
			}
			d[pos+1] = v
		}
	}
	return ans
	*/
	n := len(nums)
	if n == 0 {
		return 0
	}
	res := []int{nums[0]}
	for i := 1 ; i < n ; i++ {
		if nums[i] > res[len(res)-1] {
			res = append(res,nums[i])
		}else{
			idx := sort.Search(len(res),func(j int) bool {return res[j] >= nums[i]})
			res[idx] = nums[i]
		}
	}
	return len(res)
}
// @lc code=end

