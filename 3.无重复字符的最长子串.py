#
# @lc app=leetcode.cn id=3 lang=python3
#
# [3] 无重复字符的最长子串
#

# @lc code=start
class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        if not s:
            return 0
        
        visited = set()

        l, r = 0,0
        visited.add(s[0])

        result = 1
        for r in range(1, len(s)):
            while s[r] in visited:
                visited.remove(s[l])
                l += 1
            visited.add(s[r])
            result = max(result, len(visited))
        
        return result
            
# @lc code=end

