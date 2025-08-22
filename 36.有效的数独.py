#
# @lc app=leetcode.cn id=36 lang=python3
#
# [36] 有效的数独
#

from typing import List

# @lc code=start
class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        # 不可以用[[None]*9]*9
        rs,cs,bs = [[None for _ in range(9)] for _ in range(9)],[[None for _ in range(9)] for _ in range(9)],[[None for _ in range(9)] for _ in range(9)]
        for i in range(9):
            for j in range(9):
                item = board[i][j]
                if item not in '0123456789':
                    continue
                item = int(item) - 1
                if rs[i][item]:
                    # print(1, i, j, rs)
                    return False
                rs[i][item] = True

                if cs[j][item]:
                    # print(2, i, j)
                    return False
                cs[j][item] = True

                if bs[i // 3 * 3 + j // 3][item]:
                    # print(3, i, j)
                    return False
                bs[i // 3 * 3 + j // 3][item] = True
        return True
    
if __name__ == '__main__':
    board = [["5","3",".",".","7",".",".",".","."]
            ,["6",".",".","1","9","5",".",".","."]
            ,[".","9","8",".",".",".",".","6","."]
            ,["8",".",".",".","6",".",".",".","3"]
            ,["4",".",".","8",".","3",".",".","1"]
            ,["7",".",".",".","2",".",".",".","6"]
            ,[".","6",".",".",".",".","2","8","."]
            ,[".",".",".","4","1","9",".",".","5"]
            ,[".",".",".",".","8",".",".","7","9"]]
    s = Solution().isValidSudoku(board)

# @lc code=end

