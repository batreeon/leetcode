/*
 * @lc app=leetcode.cn id=1707 lang=golang
 *
 * [1707] 与数组中元素的最大异或值
 */

// @lc code=start
package main
import "sort"

// 因为我们对queries重新排了序，为了最终答案的顺序正确，需要记录排序时的坐标映射
// idexreflect元素值是原始下标，下标是新的下标
var idexreflect []int
// 按照query[1]进行排序
type sortQueries [][]int
func (sq sortQueries) Len() int {
	return len(sq)
}
func (sq sortQueries) Less(i,j int) bool {
	return sq[i][1] < sq[j][1]
}
func (sq sortQueries) Swap(i,j int) {
	sq[i],sq[j] = sq[j],sq[i]
	idexreflect[i],idexreflect[j] = idexreflect[j],idexreflect[i]
}
func maximizeXor(nums []int, queries [][]int) []int {
	const (
		N int = 100010 * 30
	)
	son := make([][2]int,N)
	var idx int
	insert := func(x int) {
		p := 0
		for i := 30 ; i >= 0 ; i-- {
			s := &son[p][x >> i & 1]
			if *s == 0 {
				idx++
				*s = idx
			}
			p = *s
		}
	}
	queryfunc := func(x int) int {
		res := 0
		p := 0
		for i := 30 ; i >= 0 ; i-- {
			s := x >> i & 1
			sidx := 0
			if s == 0 {
				sidx = 1
			}
			if son[p][sidx] != 0 {
				res += 1 << i
				p = son[p][sidx]
			}else{
				p = son[p][s]
			}
		}
		return res
	}
	sort.Ints(nums)
	idexreflect = make([]int,len(queries))
	for i := range idexreflect {
		idexreflect[i] = i
	}
	sq := sortQueries(queries)
	sort.Sort(sq)
	lastright := 0
	result := make([]int,len(queries))
	for i,query := range queries {
		right := sort.Search(len(nums),func(i int) bool {return nums[i] > query[1]})
		if right == 0 {
			// 没有比query[1]小的
			result[i] = -1
		}else{
			// 把不超过query[1]的都insert
			for j := lastright ; j < right ; j++ {
				insert(nums[j])
			}
			result[i] = queryfunc(query[0])
			lastright = right
		}
	}
	res := make([]int,len(queries))
	for i := range idexreflect {
		res[idexreflect[i]] = result[i]
	}
	return res
}
// @lc code=end

