#
# @lc app=leetcode.cn id=14 lang=python3
#
# [14] 最长公共前缀
#

# @lc code=start
class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        result: str = ''
        i = 0
        while True:
            if i >= len(strs[0]):
                break
            for j in range(1, len(strs)):
                if i >= len(strs[j]):
                    return result
                if strs[j][i] != strs[j-1][i]:
                    return result
            result += strs[0][i]
            i += 1

        return result

            
# @lc code=end

