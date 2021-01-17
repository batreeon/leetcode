/*
 * @lc app=leetcode.cn id=1122 lang=golang
 *
 * [1122] 数组的相对排序
 */

// @lc code=start

// type intsort []int
// func (a intsort) Len() int {
// 	return len(a)
// }
// func (a intsort) Swap(i,j int) {
// 	a[i],a[j] = a[j],a[i]
// }
// func (a intsort) Less(i,j int) bool {
// 	return a[i] < a[j]
// }
func relativeSortArray(arr1 []int, arr2 []int) []int {
	if len(arr1) == 0 {
		return arr1
	}
	if len(arr2) == 0 {
		sort.Ints(arr1)
		return arr1
	}
	j1 := 0
	for i := 0 ; i < len(arr2);i++ {
		for j2 := j1;j2 < len(arr1);j2++ {
			if arr1[j2] == arr2[i] {
				arr1[j2],arr1[j1] = arr1[j1],arr1[j2]
				j1++
			}
		}
	}
	sort.Ints(arr1[j1:])
	return arr1
}
// @lc code=end

