/*
给你一个 n x n 的整数方阵 matrix 。你可以执行以下操作 任意次 ：

选择 matrix 中 相邻 两个元素，并将它们都 乘以 -1 。
如果两个元素有 公共边 ，那么它们就是 相邻 的。

你的目的是 最大化 方阵元素的和。请你在执行以上操作之后，返回方阵的 最大 和。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-matrix-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func maxMatrixSum(matrix [][]int) int64 {
    minusCount := 0
    minAbs := abs(matrix[0][0])
    r,c := len(matrix),len(matrix[0])
    sum := 0
    for i := 0 ; i < r ; i++ {
        for j := 0 ; j < c ; j++ {
            if abs(matrix[i][j]) < minAbs {
                minAbs = abs(matrix[i][j])
            }
            sum += abs(matrix[i][j])
            if matrix[i][j] < 0 {
                minusCount = 1 - minusCount
            }
        }
    }
    if minusCount == 0 {
        return int64(sum)
    }
    sum -= minAbs*2
    return int64(sum)
}
func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}