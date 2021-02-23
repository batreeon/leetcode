/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
package main
func intersect(nums1 []int, nums2 []int) []int {
	n1,n2 := len(nums1),len(nums2)
	// if n1 == 0 || n2 == 0 {
	// 	return []int{}
	// }
	var m map[int][]int
	if n1 > n2 {
		m = make(map[int][]int,n2)//这里就别用:=,要不就成了新局部变量了
		for _,v := range nums2 {
			if _,ok := m[v] ; !ok {
				m[v] = []int{0,0}
			}
			m[v][0]++
		}
		for _,v := range nums1 {
			if _,ok := m[v] ; ok {
				m[v][1]++
			}
		}
	}else{
		m = make(map[int][]int,n1)
		for _,v := range nums1 {
			if _,ok := m[v] ; !ok {
				m[v] = []int{0,0}
			}
			m[v][0]++
		}
		for _,v := range nums2 {
			if _,ok := m[v] ; ok {
				m[v][1]++
			}
		}
	}
	ans := []int{}
	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}
	for k,v := range m {
		n := min(v[0],v[1])
		for i := 0 ; i < n ; i++ {
			ans = append(ans,k)
		}
	}
	return ans
}
// @lc code=end

