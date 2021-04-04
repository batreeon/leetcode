package main

import (
	"math"
	"strconv"
)

/*
给你一个数组 nums ，数组中只包含非负整数。定义 rev(x) 的值为将整数 x 各个数字位反转得到的结果。比方说 rev(123) = 321 ， rev(120) = 21 。我们称满足下面条件的下标对 (i, j) 是 好的 ：

0 <= i < j < nums.length
nums[i] + rev(nums[j]) == nums[j] + rev(nums[i])
请你返回好下标对的数目。由于结果可能会很大，请将结果对 109 + 7 取余 后返回。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-nice-pairs-in-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 思路：
// nums[i] + rev(nums[j]) == nums[j] + rev(nums[i])
// 即nums[i] - rev(nums[i]) == nums[j] - rev(nums[j])
// 所以只要某个数自身与自身的翻转之差相等,就可以构成一对好对子
// 就相当于握手问题,两两握手,求出所有的自身与自身翻转之差的情况,统计,然后每一种就计算

func countNicePairs(nums []int) int {
	cnt := map[int]int{}
	reverse := func(s string) string {
		ss := []byte(s)
		for i,j := 0,len(s)-1 ; i < j ; i,j = i+1,j-1 {
			ss[i],ss[j] = ss[j],ss[i]
		}
		return string(ss)
	}
	for _,num := range nums {
		b := strconv.Itoa(num)
		b = reverse(b)
		bnum,_ := strconv.Atoi(b)
		cnt[num-bnum]++
	}
	result := 0
	mod := int64(math.Pow10(9)+7)
	for _,v := range cnt {
		result = int(int64(result+v*(v-1)/2)%mod)
	}
	return result
}