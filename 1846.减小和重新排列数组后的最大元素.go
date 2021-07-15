/*
 * @lc app=leetcode.cn id=1846 lang=golang
 *
 * [1846] 减小和重新排列数组后的最大元素
 */

// @lc code=start
package main
import "sort"
func maximumElementAfterDecrementingAndRearranging(arr []int) int {
    if len(arr) == 0 {
        return arr[0]
    }
    sort.Ints(arr)
    arr[0] = 1
    for i := 1 ; i < len(arr) ; i++ {
        if arr[i] - arr[i-1] <= 1 {
            continue
        }
        arr[i] = arr[i-1]+1
    }
    return arr[len(arr)-1]
}
// @lc code=end

