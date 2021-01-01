/*
 * @lc app=leetcode.cn id=605 lang=golang
 *
 * [605] 种花问题
 */

// @lc code=start
func canPlaceFlowers(flowerbed []int, n int) bool {
	begin := -1 //无意义
	var v int
	for i , j := range flowerbed {
		if j == 1 {
			if begin == -1 {
				begin = i
				v += begin/2	//[(begin-1)-0+1] / 2
			} else{
				v += (i-begin-2)/2 //[(i-1)-begin-2 + 1] / 2
				begin = i
			}
		}
	}
	if begin == -1 {//全为0
		v = (len(flowerbed)+1)/2
	}else{
		v += (len(flowerbed)-1-begin)/2
	}
	return v >= n 
}
// @lc code=end

