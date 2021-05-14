/*
 * @lc app=leetcode.cn id=1024 lang=golang
 *
 * [1024] 视频拼接
 */

// @lc code=start
package main
import "sort"

// 对视频片段排序，按开始时间排，若开始时间相同就只看最长的
type sortclips [][]int
func (s sortclips) Len() int {
	return len(s)
}
func (s sortclips) Less(i,j int) bool {
	return s[i][0] < s[j][0] || (s[i][0] == s[j][0] && s[i][1] > s[j][1])
}
func (s sortclips) Swap(i,j int) {
	s[i],s[j] = s[j],s[i]
}

func videoStitching(clips [][]int, T int) int {
	if T == 0 {
		return 0
	}

	// pieces[i]标记剪辑到第i秒需要的片段数
	pieces := make([]int,T+1)
	s := sortclips(clips)
	sort.Sort(s)
	// 不能开始，若T为0除外，我们前面已经考虑过了
	if s[0][0] != 0 {
		return -1
	}
	// 一步到位
	if s[0][1] >= T {
		return 1
	}

	// 第一个片段
	cnt := 1
	for i := s[0][0] ; i <= s[0][1] ; i++ {
		pieces[i] = cnt
	}
	for i := 1 ; i < len(s) ; i++ {
		// 但前开始时间
		if s[i][0] == s[i-1][0] {
			continue
		}

		// 没交叉。必须要有交叉，搭上边都不行
		if pieces[s[i][0]] == 0 || pieces[s[i][0]-1] == 0 {
			return -1
		}
		// 在上一步的基础上增加一个片段
		cnt = pieces[s[i][0]]+1
		for j := s[i][1] ; j >= s[i][0] ; j-- {
			// 视频片段可能会超过T，避免越界
			if j > T {
				continue
			}
			// 可以结束了，前面的都已经被覆盖过了
			if pieces[j] != 0 {
				break
			}
			pieces[j] = cnt
		}
		// 找到了答案
		if pieces[T] != 0 {
			return pieces[T]
		}
	}
	// 没有覆盖到pieces[T]
	return -1
}
// @lc code=end

