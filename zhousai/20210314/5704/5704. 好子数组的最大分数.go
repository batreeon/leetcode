package main

/*
给你一个整数数组 nums （下标从 0 开始）和一个整数 k 。

一个子数组 (i, j) 的 分数 定义为 min(nums[i], nums[i+1], ..., nums[j]) * (j - i + 1) 。一个 好 子数组的两个端点下标需要满足 i <= k <= j 。

请你返回 好 子数组的最大可能 分数 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-score-of-a-good-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func maximumScore(nums []int, k int) int {
	/*
	n := len(nums)
	// minV := nums[k]
	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}
	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}

	// minV用来记录[i,j]子数组中的最小元素
	minV := make([][]int,n)
	for i := 0 ; i < n ; i++ {
		minV[i] = make([]int,n)
		minV[i][i] = nums[i]
	}
	for i := n-2 ; i >= 0 ; i-- {
		for j := i + 1 ; j < n ; j++ {
			minV[i][j] = min(minV[i][j-1],nums[j])
		}
	}

	ans := nums[k]
	for i := 2 ; i <= n ; i++ {
		left := max(0,k-i+1)
		right := left+i-1
		curLenMaxV := minV[left][right]
		for right < n && left <= k {
			curLenMaxV = max(curLenMaxV,minV[left][right])
			left++
			right++
		}
		ans = max(ans,i*curLenMaxV)
	}
	return ans
	*/
	max := func(x,y int) int {
		if x > y {
			return x
		}
		return y
	}
	l,r,n,res := k,k,len(nums),0
	for {
		for l >= 0 && nums[l] >= nums[k] {l--}
		for r < n && nums[r] >= nums[k] {r++}
		res = max(res,(r-l-1)*nums[k])
		if l >= 0 && r < n {
			//为什么让nums[k]等于较大的那个，一步一步来嘛，nums[k]是一直会减小的，那肯定先取较大的啊
			nums[k] = max(nums[l],nums[r])
		}else if l < 0 && r >= n {
			break
		}else if l >= 0 {
			nums[k] = nums[l]
		}else{
			nums[k] = nums[r]
		}
	}
	return res
}