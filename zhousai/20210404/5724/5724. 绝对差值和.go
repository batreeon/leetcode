package main

/*
给你两个正整数数组 nums1 和 nums2 ，数组的长度都是 n 。

数组 nums1 和 nums2 的 绝对差值和 定义为所有 |nums1[i] - nums2[i]|（0 <= i < n）的 总和（下标从 0 开始）。

你可以选用 nums1 中的 任意一个 元素来替换 nums1 中的 至多 一个元素，以 最小化 绝对差值和。

在替换数组 nums1 中最多一个元素 之后 ，返回最小绝对差值和。因为答案可能很大，所以需要对 109 + 7 取余 后返回。

|x| 定义为：

如果 x >= 0 ，值为 x ，或者
如果 x <= 0 ，值为 -x
*/
import (
	// "math"
	"math"
	"sort"
)

// 对nums1进行排序，但还要记录排序后每个元素原来的下标，因为还要根据原下标访问对应的nums2[i]
// 这一步其实有些多余，可以在交换nums1[i],nums1[j]时，同时交换nums2[i],num2[j],
// 这样就不用维护源下标了
type sortedIndex struct {
	nums  []int
	index []int
}

func (a sortedIndex) Len() int { return len(a.nums) }
func (a sortedIndex) Swap(i, j int) {
	a.index[i], a.index[j] = a.index[j], a.index[i]
	a.nums[i], a.nums[j] = a.nums[j], a.nums[i]
}
// 根据nums1中的元素值进行排序
func (a sortedIndex) Less(i, j int) bool { return a.nums[i] < a.nums[j] }

// 思路：先算出当前的差值和，然后找出nums1中替换每一位后能够得到的最大的下降
// 再找出全局最大下降，减去前面的原差值和，就是最后的结果
// 怎么找出nums1中每一个下标i对应的最大下降呢，策略用将nums1排序，利用查找算法（如二分），
// 找到对应的nums2[i]再排序后的nums1中的位置
// 此时相应位置左右两侧与nums2[i]的插值就是i这一下标处能够得到的最小差值。维护一个全局最大下降，每次更新。
// 当处理完所有下标，就可以找到最大下降，再减去原来的差值和，就行了
func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	mod := int64(math.Pow10(9)+7)
	n := len(nums1)
	diff := make([]int, n)
	abs := func(x int) int {
		if x < 0 {
			return -1 * x
		}
		return x
	}
	// 求得原始的差值和
	sumdiff := 0
	for i, num1 := range nums1 {
		diff[i] = abs(num1 - nums2[i])
		// 每加一次都求一次余
		sumdiff = int(int64(sumdiff + diff[i])%mod)
	}
	index := make([]int, n)
	for i, _ := range index {
		index[i] = i
	}

	// 初始化sortedIndex
	nums1Copy := make([]int,n)
	copy(nums1Copy,nums1)
	sortedindex := sortedIndex{
		nums:  nums1Copy,
		index: index,
	}
	sort.Sort(sortedindex)
	if sortedindex.nums[0] == sortedindex.nums[n-1] {
		return sumdiff
	}

	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}

	// 记录全局最大的下降
	maxChange := 0
	// 遍历每一位nums[i]，找出该位值的最大下降，并维护maxChange
	for i := 0 ; i < n ; i++ {
		// 没有下降空间了
		if diff[sortedindex.index[i]] == 0 {
			continue
		}
		// 找出nums2[i]在排序后的nums1中的位置，标准库使用了二分查找
		nearindex := sort.Search(n,func(ii int) bool {
			return nums2[sortedindex.index[i]] <= sortedindex.nums[ii]
		})
		if nearindex == n {
			// nums2[i]比nums1所有值都大
			change := diff[sortedindex.index[i]] - min(diff[sortedindex.index[i]],abs(nums2[sortedindex.index[i]]-sortedindex.nums[n-1]))
			if change > maxChange {
				maxChange = change
			}
		}else if nearindex == 0 {
			// nums2[i]比nums1所有值都小
			change := diff[sortedindex.index[i]] - min(diff[sortedindex.index[i]],abs(nums2[sortedindex.index[i]]-sortedindex.nums[0]))
			if change > maxChange {
				maxChange = change
			}
		}else {
			change := diff[sortedindex.index[i]] - min(diff[sortedindex.index[i]],min(abs(nums2[sortedindex.index[i]]-sortedindex.nums[nearindex]),abs(nums2[sortedindex.index[i]]-sortedindex.nums[nearindex-1])))
			if change > maxChange {
				maxChange = change
			}
		}
	}
	return sumdiff-maxChange
}
