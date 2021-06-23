func exist(board [][]byte, word string) bool {
	oris := make([][]int,4)
	oris[0] = []int{-1,0}
	oris[1] = []int{0,-1}
	oris[2] = []int{0,1}
	oris[3] = []int{1,0}
	r,c := len(board),len(board[0])
	visited := make([]bool,r*c)
	var e func(pos []int,idx int,visited []bool) bool
	e = func(pos []int,idx int,visited []bool) bool {
		if idx == len(word) {
			return true
		}
		if pos[0] < 0 || pos[0] >= r || pos[1] < 0 || pos[1] >= c || visited[pos[0]*c+pos[1]] {
			return false
		}
		if board[pos[0]][pos[1]] == word[idx] {
			visited[pos[0]*c+pos[1]] = true
			for _,ori := range oris {
				if e([]int{pos[0]+ori[0],pos[1]+ori[1]},idx+1,visited) {
					return true
				}
			}
			delete(visitd,pos[0]*c+pos[1])
		}
		return false
	}
	for i := 0 ; i < len(board) ; i++ {
		for j := 0 ; j < len(board[i]) ; j++ {
			if e([]int{i,j},0,visited) {
				return true
			}
		}
	}
	return false
}