#
# @lc app=leetcode.cn id=224 lang=python3
#
# [224] 基本计算器
#

# @lc code=start
class Solution:
    def calculate(self, s: str) -> int:
        ops = [1]
        i = 0
        result = 0
        sign = 1
        while i < len(s):
            c = s[i]
            if c == ' ':
                i += 1
            elif c == '+':
                sign = ops[-1]
                i += 1
            elif c == '-':
                sign = -ops[-1]
                i += 1
            elif c == '(':
                ops.append(sign)
                i += 1
            elif c == ')':
                ops.pop()
                i += 1
            else:
                num = 0
                while i < len(s) and s[i] in '0123456789':
                    num = num * 10 + int(s[i])
                    i += 1
                num = sign * num
                result += num
        return result
# @lc code=end

