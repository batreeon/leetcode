#
# @lc app=leetcode.cn id=1328 lang=python3
#
# [1328] 破坏回文串
#

# @lc code=start
class Solution:
    def breakPalindrome(self, palindrome: str) -> str:
        if len(palindrome) == 1:
            return ""
        for i in range(len(palindrome)):
            # 如果是正中间的元素，那么即使换了，结果依然是回文的，跳过
            if len(palindrome) % 2 == 1 and i == len(palindrome)//2:
                continue
            elif palindrome[i] != 'a':
                return palindrome[:i] + 'a' + palindrome[i+1:]
        return palindrome[:len(palindrome)-1] + 'b'
# @lc code=end

