package main
func maxValue(grid [][]int) int {
	r,c := len(grid),len(grid[0])
	f := make([]int,c+1)
	for i := range grid[0] {
		f[i+1] = f[i] + grid[0][i]
	}
	max := func(x,y int) int {
		if x > y { return x}
		return y
	}
	for i := 1 ; i < r ; i++ {
		ff := make([]int,c+1)
		for j := 0 ; j < c ; j++ {
			ff[j+1] = max(f[j+1],ff[j])+grid[i][j]
		}
		f = ff
	}
	return f[c]
}