/*
 * @lc app=leetcode.cn id=1109 lang=golang
 *
 * [1109] 航班预订统计
 */

// @lc code=start
func corpFlightBookings(bookings [][]int, n int) []int {
	result := make([]int,n)
	for _, book := range bookings {
		beg,end,seats := book[0],book[1],book[2]
		for i := beg-1 ; i < end ; i++ {
			result[i] += seats
		}
	}
	
	return result
}
// @lc code=end

