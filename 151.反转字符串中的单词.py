#
# @lc app=leetcode.cn id=151 lang=python3
#
# [151] 反转字符串中的单词
#

# @lc code=start
class Solution:
    def reverseWords(self, s: str) -> str:
        s: str = s.strip()
        i = j = len(s) - 1
        result:list[str] = []
        while i >= 0:
            while i >= 0 and s[i] != ' ':
                i -= 1
            result.append(s[i+1:j+1])
            while i >= 0 and s[i] == ' ':
                i -= 1
            j = i
        return ' '.join(result)
# @lc code=end

