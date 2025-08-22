#
# @lc app=leetcode.cn id=54 lang=python3
#
# [54] 螺旋矩阵
#

# @lc code=start
class Solution:
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        dirs = [[0,1], [1,0], [0,-1], [-1, 0]]
        result = []
        i = j = di = 0
        r, c = len(matrix), len(matrix[0])
        while len(result) < r * c:
            result.append(matrix[i][j])
            matrix[i][j] = None
            nexti, nextj = i + dirs[di][0], j + dirs[di][1]

            if nexti < 0 or nexti == r or nextj < 0 or nextj == c or matrix[nexti][nextj] is None:
                di = (di + 1) % 4
            i ,j = i + dirs[di][0], j + dirs[di][1]
        return result

# @lc code=end

