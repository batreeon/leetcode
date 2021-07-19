/*
 * @lc app=leetcode.cn id=1838 lang=golang
 *
 * [1838] 最高频元素的频数
 */

// @lc code=start
package main
import "sort"

func maxFrequency(nums []int, k int) int {
	/*
	// 哎，不知道错哪儿了
	sort.Ints(nums)
	presum := make([]int,len(nums))
	presum[0] = nums[0]
	for i := 1 ; i < len(nums) ; i++ {
		presum[i] = presum[i-1] + nums[i]
	}
	// 这个错误的原因之一：不一定从第一个就开始加啊，也就是k次增1所作用的数据区间不一定从下标0开始啊
	idx := sort.Search(len(nums),func(i int) bool {return presum[i]+k < nums[i]*(i+1)})
	return idx
	maxLen := idx
	for len(nums) >= maxLen {
		for i := range presum {
			presum[i] -= presum[0]
		}
		nums = nums[1:]
		idx := sort.Search(len(nums),func(i int) bool {return presum[i]+k < nums[i]*(i+1)})
		if idx > maxLen {
			maxLen = idx
		}
	}
	return maxLen
	*/

	
	// 下面这个也不行，我吐了
	sort.Ints(nums)
	presum := make([]int,len(nums)+1)
	for i := 1 ; i <= len(nums) ; i++ {
		presum[i] = presum[i-1] + nums[i-1]
	}
	l,r := 0,len(nums)
	check := func(length int) bool {
		for i := 0 ; i + length - 1 < len(nums) ; i++ {
			r := i+length-1
			cur := presum[r+1]-presum[i]
			total := nums[r]*length
			if total-cur <= k {
				return true
			}
		}
		return false
	}
	for l < r {
		mid := (l+r+1) >> 1
		if check(mid) {
			l = mid
		}else{
			r = mid-1
		}
	}

	return r
	

	/*
	// 成了
	sort.Ints(nums)
	result := 1
	for l,r,total := 0,1,0 ; r < len(nums) ; r++ {
		total += (nums[r]-nums[r-1]) * (r-l)
		for total > k {
			total -= nums[r]-nums[l]
			l++
		}
		if r-l+1 > result {
			result = r-l+1
		}
	}
	return result
	*/
}
// @lc code=end

