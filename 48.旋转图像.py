#
# @lc app=leetcode.cn id=48 lang=python3
#
# [48] 旋转图像
#

# @lc code=start
class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        n = len(matrix)
        for i in range(n//2):
            for j in range(i, n-i-1):
                i1, j1 = i, j
                i2, j2 = self.get_rotated(i1, j1, n)
                i3, j3 = self.get_rotated(i2, j2, n)
                i4, j4 = self.get_rotated(i3, j3, n)
                matrix[i1][j1], matrix[i2][j2], matrix[i3][j3], matrix[i4][j4] = \
                matrix[i2][j2], matrix[i3][j3], matrix[i4][j4],matrix[i1][j1]

    def get_rotated(self, i, j, n):
        return n-j-1, i
# @lc code=end

