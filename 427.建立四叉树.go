/*
 * @lc app=leetcode.cn id=427 lang=golang
 *
 * [427] 建立四叉树
 */

// @lc code=start
/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func construct(grid [][]int) *Node {
    n := len(grid)
	sumgrid := make([][]int, n)
	for i := 0; i < n; i++ {
		sumgrid[i] = make([]int, n)
	}
	sumgrid[0][0] = grid[0][0]
	for j := 1; j < n; j++ {
		sumgrid[0][j] = sumgrid[0][j-1] + grid[0][j]
	}
	for i := 1; i < n; i++ {
		sumgrid[i][0] = sumgrid[i-1][0] + grid[i][0]
	}
	for i := 1; i < n; i++ {
		for j:= 1; j< n; j++ {
			sumgrid[i][j] = sumgrid[i-1][j] + sumgrid[i][j-1] + grid[i][j] - sumgrid[i-1][j-1]
		}
	}
	return bt(grid, sumgrid, 0, 0, n)
}

func bt(grid [][]int, sumgrid [][]int, i, j, n int) *Node {
	if n == 1 {
		return &Node{
			IsLeaf: true,
			Val: grid[i][j] == 1,
		}
	}
	if i == 0 && j == 0 {
		sum := sumgrid[i+n-1][j+n-1]
		if sum == n*n || sum == 0{
			return &Node{
				IsLeaf: true,
				Val: sum != 0,
			}
		}
	}else if i == 0 {
		sum := sumgrid[i+n-1][j+n-1] - sumgrid[i+n-1][j-1]
		if sum == n*n || sum == 0{
			return &Node{
				IsLeaf: true,
				Val: sum != 0,
			}
		}
	}else if j == 0 {
		sum := sumgrid[i+n-1][j+n-1] - sumgrid[i-1][j+n-1]
		if sum == n*n || sum == 0{
			return &Node{
				IsLeaf: true,
				Val: sum != 0,
			}
		}
	}else{
		sum := sumgrid[i+n-1][j+n-1] - sumgrid[i-1][j+n-1] - sumgrid[i+n-1][j-1] + sumgrid[i-1][j-1]
		if sum == n*n || sum == 0{
			return &Node{
				IsLeaf: true,
				Val: sum != 0,
			}
		}
	}
	return &Node{
		IsLeaf: false,
		TopLeft: bt(grid, sumgrid, i, j, n/2),
		TopRight: bt(grid, sumgrid, i, j+n/2, n/2),
		BottomLeft: bt(grid, sumgrid, i + n/2, j, n/2),
		BottomRight: bt(grid, sumgrid, i + n/2, j + n/2, n/2),
	}
}
// @lc code=end

