package main
func movingCount(m int, n int, k int) int {
	visited := make([]bool,m*n)

	var mc func(x,y int) int
	var check func(x,y int) bool
	var sum func(x int) int

	mc = func(x,y int) int {
		if check(x,y) {
			visited[x*n+m] = true
			return 1+mc(x-1,y)+mc(x,y-1)+mc(x,y+1)+mc(x+1,y)
		}
		return 0
	}

	check = func(x,y int) bool {
		return !( x < 0 || x >= m || y < 0 || y >= n || k > sum(x)+sum(y) || visited[x*n+m])
	}

	sum = func(x int) int {
		result := 0
		for x > 0 {
			result += x % 10
			x /= 10
		}
		return result
	}

	return mc(0,0)
}