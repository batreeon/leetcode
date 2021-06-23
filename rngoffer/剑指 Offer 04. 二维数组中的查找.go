func findNumberIn2DArray(matrix [][]int, target int) bool {
	r,c := len(matrix),len(matrix[0])
	for i := 0 ; i < r ; i++ {
		for j := c ; j >= 0 ; j-- {
			if matrix[i][j] == target {
				return true
			}
			if matrix[i][j] > target {
				// 右下角全是比当前matrix[i][j]大的，就不用看了
				c = j-1
			}
			if matrix[i][j] < target {
				// j是递减考察的，这一行左边都更小，没必要看了，直接看下一层
				break
			}
		}
	}
	return false
}