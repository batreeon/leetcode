#
# @lc app=leetcode.cn id=58 lang=python3
#
# [58] 最后一个单词的长度
#

# @lc code=start
class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        result: int = 0
        l = 0
        for a in s:
            if a != ' ':
                l += 1
            if a == ' ':
                if l > 0:
                    result = l
                l = 0
        if l > 0:
            result = l

        return result
# @lc code=end

