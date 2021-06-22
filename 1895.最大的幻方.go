/*
 * @lc app=leetcode.cn id=1895 lang=golang
 *
 * [1895] 最大的幻方
 */

// @lc code=start
func largestMagicSquare(grid [][]int) int {
	result := 1
	min := func(x,y int) int {
		if x < y {
			return x
		}
		return y
	}
	m,n := len(grid),len(grid[0])
	for i := 2 ; i <= min(m,n) ; i++ {
		nexti := false
		for top := 0 ; top+i-1 < m ; top++ {
			bot := top+i-1
			for left := 0 ; left+i-1 < n ; left++ {
				next := false
				right := left+i-1
				sum := 0
				for y := left ; y <= right ; y++ {
					sum += grid[top][y]
				}
				// 行
				for x := top+1 ; x <= bot ; x++ {
					curs := 0
					for y := left ; y <= right ; y++ {
						curs += grid[x][y]
					}
					if curs != sum {
						next = true
						break
					}
				}
				if next {
					continue
				}
				// 列
				for y := left ; y <= right ; y++ {
					curs := 0
					for x := top ; x <= bot ; x++ {
						curs += grid[x][y]
					}
					if curs != sum {
						next = true
						break
					}
				}
				if next {
					continue
				}
				// 对角线
				curs := 0
				for x,y := top,left ; x <= bot ; x,y = x+1,y+1 {
					curs += grid[x][y]
				}
				if curs != sum {
					continue
				}
				curs = 0
				for x,y := top,right ; x <= bot ; x,y = x+1,y-1 {
					curs += grid[x][y]
				}
				if curs != sum {
					continue
				}
				if i > result {
					result = i
					nexti = true
				}
				if nexti {
					break
				}
			}
			if nexti {
				break
			}
		}
	}
	return result
}
// @lc code=end

