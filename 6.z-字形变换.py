#
# @lc app=leetcode.cn id=6 lang=python3
#
# [6] Z 字形变换
#

# @lc code=start
class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows == 1:
            return s
        rows: list[str] = []
        for i in range(min(len(s), numRows)):
            rows.append('')
        
        down = False
        curRow = 0
        for c in s:
            rows[curRow] += c
            if curRow == 0 or curRow == numRows - 1:
                down = not down
            curRow += 1 if down else -1
        result = ''.join(rows)
        return result
# @lc code=end

