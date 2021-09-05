func findFarmland(land [][]int) [][]int {
	r,c := len(land),len(land[0])
	result := [][]int{}
	for i := 0 ; i < r ; i++ {
		for j := 0 ; j < c ; j++ {
			if land[i][j] == 1 {
				bottom,right := r-1,c-1
				for x := i ; x <= bottom ; x++ {
					if land[x][j] == 0 {
						bottom = x-1
						break
					}
					for y := j ; y <= right ; y++ {
						if land[x][y] == 1 {
							land[x][y] = 0
						}else{
							right = y-1
						}
					}
				}
                result = append(result,[]int{i,j,bottom,right})
			}
		}
	}
	return result
}