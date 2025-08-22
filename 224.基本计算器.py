#
# @lc app=leetcode.cn id=224 lang=python3
#
# [224] 基本计算器
#

# @lc code=start
#  只有+和-，那么可以括号外的单个数字看作一个整体、把括号看作一个整体
# ops里记录每块整体的符号，sign记录每个数字的符号，每个数字的符号需要乘以每个整体的符号
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

